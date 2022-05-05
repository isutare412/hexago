package donation

import (
	"context"
	"fmt"

	"github.com/isutare412/hexago/payment/pkg/port"
)

type Service struct {
	dnxRepo port.DonationRepository
}

func NewService(
	dnxRepo port.DonationRepository,
) *Service {
	return &Service{
		dnxRepo: dnxRepo,
	}
}

func (s *Service) Donate(
	ctx context.Context,
	donatorId, donateeId string,
	cents int64,
) error {
	err := s.dnxRepo.RecordDonationHistory(ctx, donatorId, donateeId, cents)
	if err != nil {
		return fmt.Errorf("recording donation history: %w", err)
	}
	return nil
}
