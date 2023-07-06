package entities

import (
	"errors"
	"fmt"
	"net/mail"
	"strconv"
	"time"

	"github.com/Nhanderu/brdoc"
)

var (
	ErrInvalidStudentID = errors.New("invalid student id")
	ErrInvalidCPF       = errors.New("invalid cpf")
	ErrInvalidEmail     = errors.New("invalid email")
	ErrInvalidBirthDate = errors.New("invalid birth date")
)

type Student struct {
	ID        string
	Name      string
	CPF       string
	Email     string
	BirthDate time.Time
}

func NewStudent(id string, name string, cpf string, email string, birthDate string) (Student, error) {
	if _, err := strconv.Atoi(id); err != nil {
		return Student{}, fmt.Errorf("%w: %s", ErrInvalidStudentID, err)
	}

	if _, err := strconv.Atoi(cpf); err != nil && !brdoc.IsCPF(cpf) {
		return Student{}, ErrInvalidCPF
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return Student{}, fmt.Errorf("%w: %s", ErrInvalidEmail, err)
	}

	b, err := time.Parse(time.DateOnly, birthDate)
	if err != nil {
		return Student{}, fmt.Errorf("%w: %s", ErrInvalidBirthDate, err)
	}

	return Student{
		ID:        id,
		Name:      name,
		CPF:       cpf,
		Email:     email,
		BirthDate: b,
	}, nil
}
