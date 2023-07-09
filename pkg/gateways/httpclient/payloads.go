package httpclient

import (
	"strconv"

	"github.com/tccav/student-creator-dispatcher/pkg/domain/students"
)

type GetCourseResp struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MinPeriods int    `json:"minimum_periods_qty"`
	MaxPeriods int    `json:"maximum_periods_qty"`
}

type PostStudentsReq struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
	Email     string `json:"email"`
	CourseID  string `json:"course_id"`
}

func FromRegisterInput(input students.RegisterInput) (PostStudentsReq, error) {
	idInt, err := strconv.Atoi(input.ID)
	if err != nil {
		return PostStudentsReq{}, err
	}
	return PostStudentsReq{
		ID:        idInt,
		Name:      input.Name,
		CPF:       input.CPF,
		BirthDate: input.BirthDate,
		Email:     input.Email,
		CourseID:  input.CourseID,
	}, nil
}

type ServiceErrorResp struct {
	Description string `json:"description"`
	Status      int    `json:"status"`
	Message     string `json:"message"`
}
