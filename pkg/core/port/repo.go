package port

import (
	"context"

	"github.com/isutare412/hexago/pkg/core/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Sessional interface {
	StartSession(context.Context) (mongo.SessionContext, error)
}

type StudentRepository interface {
	Sessional
	InsertStudent(ctx context.Context, stu *entity.Student) (*entity.Student, error)
	FindStudentById(ctx context.Context, id primitive.ObjectID) (*entity.Student, error)
	DeleteStudentById(ctx context.Context, id primitive.ObjectID) (int64, error)
}
