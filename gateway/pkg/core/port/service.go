package port

import (
	"context"

	centity "github.com/isutare412/hexago/common/pkg/entity"
)

type UserService interface {
	Create(ctx context.Context, user *centity.User) error
	GetById(ctx context.Context, id string) (*centity.User, error)
	DeleteById(ctx context.Context, id string) error
}

type DonationService interface {
	RequestDonation(ctx context.Context, donatorId, donateeId string, cents int64) error
}
