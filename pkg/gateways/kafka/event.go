package kafka

import "github.com/tccav/student-creator-dispatcher/pkg/domain/students"

type EventMetadata struct {
	ID   string `json:"event_id"`
	Type string `json:"event_type"`
}

type StudentRegisteredEvent struct {
	EventMetadata
	Payload StudentRegisteredPayload `json:"payload"`
}

type StudentRegisteredPayload struct {
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date"`
	CourseID  string `json:"course_id"`
}

func (p StudentRegisteredPayload) toRegisterStudentInput() students.RegisterStudentInput {
	return students.RegisterStudentInput{
		ID:        p.StudentID,
		Name:      p.Name,
		CPF:       p.CPF,
		Email:     p.Email,
		BirthDate: p.BirthDate,
		CourseID:  p.CourseID,
	}
}
