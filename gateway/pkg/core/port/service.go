package port

import (
	"context"

	"github.com/isutare412/hexago/gateway/pkg/core/entity"
)

type UserService interface {
	Create(ctx context.Context, user *entity.User) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	DeleteByEmail(ctx context.Context, email string) error
}
