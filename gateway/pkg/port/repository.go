package port

import (
	"context"

	centity "github.com/isutare412/hexago/common/pkg/entity"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *centity.User) error
	FindUserById(ctx context.Context, id string) (*centity.User, error)
	DeleteUserById(ctx context.Context, id string) error
}
