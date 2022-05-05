package main

import (
	"context"

	"github.com/isutare412/hexago/payment/pkg/config"
	"github.com/isutare412/hexago/payment/pkg/core/service/donation"
	"github.com/isutare412/hexago/payment/pkg/infrastructure/mongo"
)

type Components struct {
	mongoRepo       *mongo.Repository
	donationService *donation.Service
}

func dependencyInjection(
	ctx context.Context,
	cfg *config.Config,
) (*Components, error) {
	mongoRepo, err := mongo.NewRepository(ctx, cfg.MongoDB)
	if err != nil {
		return nil, err
	}

	donationService := donation.NewService(mongoRepo)

	return &Components{
		mongoRepo:       mongoRepo,
		donationService: donationService,
	}, nil
}
