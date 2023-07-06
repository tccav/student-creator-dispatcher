package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Configs struct {
	Telemetry   Telemetry
	Consumer    consumer
	DB          db
	Kafka       kafka
	StudentsAPI studentsAPI
	CoursesAPI  coursesAPI
}

type Telemetry struct {
	OtelCollector string `envconfig:"OTEL_URL" required:"true"`
	Environment   string `envconfig:"ENVIRONMENT" default:"dev"`
	ServiceName   string `envconfig:"SERVICE_NAME" default:"student-creator-dispatcher"`
}

type consumer struct {
	Group   string        `envconfig:"CONSUMER_GROUP" default:"student_creator_dispatcher"`
	Timeout time.Duration `envconfig:"CONSUMER_TIMEOUT" default:"15s"`
}

type db struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
	Options  string `envconfig:"DB_OPTIONS"`
}

func (d db) URL() string {
	u := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", d.User, d.Password, d.Host, d.Port, d.Name)
	if d.Options != "" {
		u += fmt.Sprintf("?%s", d.Options)
	}
	return u
}

type kafka struct {
	Host     string `envconfig:"KAFKA_HOST" required:"true"`
	Port     string `envconfig:"KAFKA_PORT" required:"true"`
	User     string `envconfig:"KAFKA_USER"`
	Password string `envconfig:"KAFKA_PASSWORD"`
}

func (k kafka) URL() string {
	return fmt.Sprintf("%s:%s", k.Host, k.Port)
}

type studentsAPI struct {
	BaseURL string `envconfig:"STUDENTS_BASE_URL" required:"true"`
}

type coursesAPI struct {
	BaseURL string `envconfig:"COURSES_BASE_URL" required:"true"`
}

func LoadConfigs() (Configs, error) {
	var config Configs
	err := envconfig.Process("", &config)
	if err != nil {
		return Configs{}, err
	}
	return config, nil
}
