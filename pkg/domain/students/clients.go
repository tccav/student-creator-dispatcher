package students

import "context"

type RegisterInput struct {
	ID        string
	Name      string
	CPF       string
	Email     string
	BirthDate string
	CourseID  string
}

type StudentRegistererClient interface {
	Register(ctx context.Context, input RegisterInput) error
}

type GetCourseOutput struct {
	ID   string
	Name string
}

type CourseListerClient interface {
	GetCourse(ctx context.Context, id string) (GetCourseOutput, error)
}
