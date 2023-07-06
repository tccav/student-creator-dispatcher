package students

import (
	"context"
	"errors"
)

//go:generate moq -out idmocks/mock_usecases.go -pkg idmocks . RegisterUseCases AuthenticationUseCases

var (
	ErrInvalidCourseID      = errors.New("invalid course id")
	ErrStudentAlreadyExists = errors.New("student already exists")
)

type RegisterStudentInput struct {
	ID        string
	Name      string
	CPF       string
	Email     string
	BirthDate string
	CourseID  string
}

type RegisterUseCases interface {
	RegisterStudent(ctx context.Context, input RegisterStudentInput) (string, error)
}
