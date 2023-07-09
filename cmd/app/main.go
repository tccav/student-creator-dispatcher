package main

import (
	"context"
	"fmt"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
	"github.com/twmb/franz-go/plugin/kotel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/tccav/student-creator-dispatcher/pkg/config"
	"github.com/tccav/student-creator-dispatcher/pkg/domain/students/stdusecases"
	"github.com/tccav/student-creator-dispatcher/pkg/gateways/httpclient"
	"github.com/tccav/student-creator-dispatcher/pkg/gateways/kafka"
	"github.com/tccav/student-creator-dispatcher/pkg/gateways/opentelemetry"
)

var (
	AppVersion = "unknown"
	GoVersion  = "unknown"
	BuildTime  = "unknown"
)

func main() {
	logConfig := zap.NewProductionConfig()
	logConfig.DisableStacktrace = true

	logger, err := logConfig.Build(
		zap.Fields(
			zap.String("version", AppVersion),
			zap.String("go_version", GoVersion),
			zap.String("build_time", BuildTime),
		),
	)
	if err != nil {
		panic(fmt.Sprintf("unable to initialize logger: %s", err))
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	logger.Info("application init started, configs will be loaded")

	configs, err := config.LoadConfigs()
	if err != nil {
		logger.Error("failed to load configs", zap.Error(err))
		return
	}

	logger = logger.With(zap.String("environment", configs.Telemetry.Environment))

	tpClose, err := opentelemetry.InitProvider(logger, AppVersion, configs.Telemetry)
	if err != nil {
		logger.Error("failed to initialize tracer", zap.Error(err))
		return
	}
	defer tpClose()

	ctx := context.Background()
	dbCfg, err := pgxpool.ParseConfig(configs.DB.URL())
	if err != nil {
		logger.Error("failed to parse db url", zap.Error(err))
	}

	dbCfg.ConnConfig.Tracer = otelpgx.NewTracer()

	pool, err := pgxpool.NewWithConfig(ctx, dbCfg)
	if err != nil {
		logger.Error("failed to start db", zap.Error(err))
		return
	}
	defer pool.Close()
	logger.Info("db conn pool fetched")

	// Create a new kotel tracer.
	tracerOpts := []kotel.TracerOpt{
		kotel.TracerProvider(otel.GetTracerProvider()),
		kotel.TracerPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{})),
	}
	tracer := kotel.NewTracer(tracerOpts...)

	// Create a new kotel service.
	kotelOps := []kotel.Opt{
		kotel.WithTracer(tracer),
	}
	kotelService := kotel.NewKotel(kotelOps...)

	kOpts := []kgo.Opt{
		kgo.SeedBrokers(configs.Kafka.URL()),
		kgo.WithHooks(kotelService.Hooks()...),
		kgo.ConsumeTopics("identity.cdc.students.0"),
		kgo.ConsumerGroup("student-creator-dispatcher"),
	}
	if configs.Kafka.User != "" {
		kOpts = append(kOpts, kgo.SASL(plain.Auth{
			User: configs.Kafka.User,
			Pass: configs.Kafka.Password,
		}.AsMechanism()))
	}

	kafkaClient, err := kgo.NewClient(kOpts...)
	if err != nil {
		logger.Error("unable to connect to kafka broker", zap.Error(err))
		return
	}
	defer kafkaClient.Close()

	err = kafkaClient.Ping(ctx)
	if err != nil {
		logger.Error("kafka broker unreachable", zap.Error(err))
		return
	}
	logger.Info("kafka client created")

	studentsClient, err := httpclient.NewStudentsClient(configs.StudentsAPI.BaseURL)
	if err != nil {
		logger.Error("failed to create students http client", zap.Error(err))
		return
	}

	coursesClient, err := httpclient.NewCoursesClient(configs.CoursesAPI.BaseURL)
	if err != nil {
		logger.Error("failed to create courses http client", zap.Error(err))
		return
	}

	useCase := stdusecases.NewRegisterUseCase(studentsClient, coursesClient)
	recordHandler := kafka.NewStudentRegisterRecordHandler(logger, useCase)

	for {
		fetches := kafkaClient.PollFetches(context.Background())
		if fetches.IsClientClosed() {
			return
		}
		fetches.EachError(func(topic string, partition int32, err error) {
			logger.Error("fetch failed",
				zap.Error(err),
				zap.String("topic", topic),
				zap.Int32("partition", partition),
			)
		})

		fetches.EachRecord(func(record *kgo.Record) {
			span := trace.SpanFromContext(record.Context)
			span.SpanContext().TraceID()

			recordLogger := logger.With(
				zap.String("topic", record.Topic),
				zap.Int32("partition", record.Partition),
				zap.String("trace_id", span.SpanContext().TraceID().String()),
			)

			recordLogger.Info("consuming incoming message")

			handlerErr := recordHandler.Handle(record)
			if handlerErr != nil {
				recordLogger.Error("failed to consume message", zap.Error(handlerErr))
				return
			}

			recordLogger.Info("successfully consumed message")
		})
	}
}
