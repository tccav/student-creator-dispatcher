package stdusecases

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/tccav/student-creator-dispatcher/pkg/domain/students"
)

type RegisterUseCase struct {
	tracer trace.Tracer

	s students.StudentRegistererClient
	c students.CourseListerClient
}

func NewRegisterUseCase(studentRegistererClient students.StudentRegistererClient, courseListerClient students.CourseListerClient) RegisterUseCase {

	return RegisterUseCase{
		tracer: otel.Tracer(tracerName),
		s:      studentRegistererClient,
		c:      courseListerClient,
	}
}

func (r RegisterUseCase) RegisterStudent(ctx context.Context, input students.RegisterStudentInput) (string, error) {
	ctx, span := r.tracer.Start(ctx, "RegisterUseCase.RegisterStudent")
	defer span.End()

	// call institutes service to verify course id
	_, err := r.c.GetCourse(ctx, input.CourseID)
	if err != nil {
		span.RecordError(err)
		return "", err
	}

	// call student service to register student
	err = r.s.Register(ctx, students.RegisterInput{
		ID:        input.ID,
		Name:      input.Name,
		CPF:       input.CPF,
		Email:     input.Email,
		BirthDate: input.BirthDate,
		CourseID:  input.CourseID,
	})
	if err != nil {
		span.RecordError(err)
		return "", err
	}

	return input.ID, nil
}
