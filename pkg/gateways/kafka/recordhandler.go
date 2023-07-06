package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/plugin/kotel"
	"go.uber.org/zap"

	"github.com/tccav/student-creator-dispatcher/pkg/domain/students"
)

type StudentRegisterRecordHandler struct {
	logger *zap.Logger
	tracer *kotel.Tracer

	useCase students.RegisterUseCases
}

func NewStudentRegisterRecordHandler(logger *zap.Logger, useCase students.RegisterUseCases) StudentRegisterRecordHandler {
	return StudentRegisterRecordHandler{
		logger:  logger,
		tracer:  kotel.NewTracer(),
		useCase: useCase,
	}
}

func (h StudentRegisterRecordHandler) Handle(record *kgo.Record) error {
	ctx, span := h.tracer.WithProcessSpan(record)
	defer span.End()

	var event StudentRegisteredEvent
	err := json.Unmarshal(record.Value, &event)
	if err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	_, err = h.useCase.RegisterStudent(ctx, event.Payload.toRegisterStudentInput())
	if err != nil {
		h.logger.Error("failed to register student", zap.Error(err))
		return nil
	}

	// TODO: produce event back informing everything ok

	return nil
}
