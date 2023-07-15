package students

import (
	"context"
	"errors"
)

var (
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
