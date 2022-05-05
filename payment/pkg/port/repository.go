package port

import (
	"context"
)

type DonationRepository interface {
	RecordDonationHistory(ctx context.Context, donatorId, donateeId string, cents int64) error
}
