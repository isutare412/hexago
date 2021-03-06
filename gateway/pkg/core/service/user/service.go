package user

import (
	"context"
	"fmt"

	centity "github.com/isutare412/hexago/common/pkg/entity"
	"github.com/isutare412/hexago/gateway/pkg/port"
)

type Service struct {
	userRepo port.UserRepository
}

func NewService(userRepo port.UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) Create(ctx context.Context, user *centity.User) error {
	user.DonatedFrom = make([]*centity.DonateRelation, 0)
	user.DonatedTo = make([]*centity.DonateRelation, 0)

	if err := s.userRepo.InsertUser(ctx, user); err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}
	return nil
}

func (s *Service) GetById(
	ctx context.Context,
	id string,
) (*centity.User, error) {
	user, err := s.userRepo.FindUserById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("finding user by id: %w", err)
	}
	return user, nil
}

func (s *Service) DeleteById(ctx context.Context, id string) error {
	if err := s.userRepo.DeleteUserById(ctx, id); err != nil {
		return fmt.Errorf("deleting user by id: %w", err)
	}
	return nil
}
