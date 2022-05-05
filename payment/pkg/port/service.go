package port

import "context"

type DonationService interface {
	Donate(ctx context.Context, donatorId, donateeId string, cents int64) error
}
