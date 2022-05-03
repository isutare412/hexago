package main

import (
	"context"

	"github.com/isutare412/hexago/gateway/pkg/config"
	"github.com/isutare412/hexago/gateway/pkg/controller/http"
	"github.com/isutare412/hexago/gateway/pkg/core/service/donation"
	"github.com/isutare412/hexago/gateway/pkg/core/service/user"
	"github.com/isutare412/hexago/gateway/pkg/infrastructure/kafka"
	"github.com/isutare412/hexago/gateway/pkg/infrastructure/mongo"
)

type components struct {
	mongoRepo        *mongo.Repository
	donationProducer *kafka.Producer
	userService      *user.Service
	donationService  *donation.Service
	httpServer       *http.Server
}

func dependencyInjection(
	ctx context.Context,
	cfg *config.Config,
) (*components, error) {
	mongoRepo, err := mongo.NewRepository(ctx, cfg.MongoDB)
	if err != nil {
		return nil, err
	}

	donationProducer, err := kafka.NewProducer(
		ctx,
		cfg.Kafka,
		cfg.Kafka.Topics.DonationRequest)
	if err != nil {
		return nil, err
	}

	userService := user.NewService(mongoRepo)
	donationService := donation.NewService(
		mongoRepo,
		donationProducer)

	httpServer := http.NewServer(
		cfg.Server.Http,
		userService,
		donationService)

	return &components{
		mongoRepo:        mongoRepo,
		donationProducer: donationProducer,
		userService:      userService,
		donationService:  donationService,
		httpServer:       httpServer,
	}, nil
}
