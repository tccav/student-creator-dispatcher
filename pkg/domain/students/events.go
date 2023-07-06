package students

import (
	"context"

	"github.com/tccav/student-creator-dispatcher/pkg/domain/entities"
)

type Producer interface {
	ProduceStudentRegistered(ctx context.Context, student entities.Student, courseID string) error
}
