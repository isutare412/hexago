package user

import (
	"context"
	"fmt"

	"github.com/isutare412/hexago/pkg/core/entity"
	"github.com/isutare412/hexago/pkg/core/port"
)

type Service struct {
	userRepo port.UserRepo
}

func NewService(userRepo port.UserRepo) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) Create(ctx context.Context, user *entity.User) error {
	if err := s.userRepo.InsertUser(ctx, user); err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}
	return nil
}

func (s *Service) GetByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("finding user by email: %w", err)
	}
	return user, nil
}

func (s *Service) DeleteByEmail(ctx context.Context, email string) error {
	if err := s.userRepo.DeleteUserByEmail(ctx, email); err != nil {
		return fmt.Errorf("deleting user by email: %w", err)
	}
	return nil
}
