package donation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/isutare412/hexago/payment/pkg/config"
	"github.com/isutare412/hexago/payment/pkg/core/service/donation"
	"github.com/isutare412/hexago/payment/pkg/infrastructure/mongo"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestDonationService(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cfg, err := loadDonationTestConfig()
	assert.NoError(t, err)

	mongoRepo, err := mongo.NewRepository(ctx, cfg.MongoDB)
	assert.NoError(t, err)

	dnxSrv := donation.NewService(mongoRepo)

	err = dnxSrv.Donate(ctx, "foo", "bar", 250)
	t.Logf("donation error: %v", err)
}

func loadDonationTestConfig() (*config.Config, error) {
	cfg, err := config.Load("../../../../configs/local/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}

	cfg.MongoDB.Database = "hexago_test"
	return cfg, nil
}
