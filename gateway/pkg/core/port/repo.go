package port

import (
	"context"

	centity "github.com/isutare412/hexago/common/pkg/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type Sessional interface {
	StartSession(context.Context) (mongo.SessionContext, error)
}

type UserRepo interface {
	Sessional
	InsertUser(ctx context.Context, user *centity.User) error
	FindUserById(ctx context.Context, id string) (*centity.User, error)
	DeleteUserById(ctx context.Context, id string) error
}
