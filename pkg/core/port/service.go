package port

import (
	"context"

	"github.com/isutare412/hexago/pkg/core/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentService interface {
	AddStudent(ctx context.Context, stu *entity.Student) (*entity.Student, error)
	StudentById(ctx context.Context, id primitive.ObjectID) (*entity.Student, error)
	RemoveStudentById(ctx context.Context, id primitive.ObjectID) error
}
