package main

import (
	"context"

	"github.com/isutare412/hexago/payment/pkg/config"
	"github.com/isutare412/hexago/payment/pkg/controller/mq"
	"github.com/isutare412/hexago/payment/pkg/core/service/donation"
	"github.com/isutare412/hexago/payment/pkg/infrastructure/kafka"
	"github.com/isutare412/hexago/payment/pkg/infrastructure/mongo"
)

type Components struct {
	mongoRepo        *mongo.Repository
	donationConsumer *kafka.Consumer
	donationService  *donation.Service
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

	donationHandler := mq.NewDonationHandler(
		cfg.Kafka.Topics.DonationRequest,
		donationService)
	donationConsumer, err := kafka.NewConsumer(
		ctx,
		cfg.Kafka,
		cfg.Kafka.Topics.DonationRequest,
		donationHandler)
	if err != nil {
		return nil, err
	}

	return &Components{
		mongoRepo:        mongoRepo,
		donationConsumer: donationConsumer,
		donationService:  donationService,
	}, nil
}
