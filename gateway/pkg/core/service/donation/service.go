package donation

import (
	"context"
	"fmt"

	pbPay "github.com/isutare412/hexago/common/pkg/pb/payment"
	"github.com/isutare412/hexago/gateway/pkg/port"
)

type Service struct {
	userRepo port.UserRepo
	payMq    port.PaymentMessageQueue
}

func NewService(
	userRepo port.UserRepo,
	payMq port.PaymentMessageQueue,
) *Service {
	return &Service{
		userRepo: userRepo,
		payMq:    payMq,
	}
}

func (s *Service) RequestDonation(
	ctx context.Context,
	donatorId, donateeId string,
	cents int64,
) error {
	_, err := s.userRepo.FindUserById(ctx, donatorId)
	if err != nil {
		return fmt.Errorf("finding donator: %w", err)
	}
	_, err = s.userRepo.FindUserById(ctx, donateeId)
	if err != nil {
		return fmt.Errorf("finding donatee: %w", err)
	}

	err = s.payMq.SendDonationRequest(&pbPay.DonationRequest{
		DonatorId: donatorId,
		DonateeId: donateeId,
		Cents:     cents,
	})
	if err != nil {
		return fmt.Errorf("sending donation request: %w", err)
	}
	return nil
}
