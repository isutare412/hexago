package port

import (
	"context"

	"github.com/isutare412/hexago/pkg/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type Sessional interface {
	StartSession(context.Context) (mongo.SessionContext, error)
}

type UserRepo interface {
	Sessional
	InsertUser(ctx context.Context, user *entity.User) error
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	DeleteUserByEmail(ctx context.Context, email string) error
}
