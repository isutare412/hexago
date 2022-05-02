package main

import (
	"context"
	"fmt"

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
	diDone := make(chan *components)
	diFail := make(chan error)
	defer close(diDone)
	defer close(diFail)

	go func() {
		mongoRepo, err := mongo.NewRepository(ctx, cfg.MongoDB)
		if err != nil {
			diFail <- err
			return
		}

		donationProducer, err := kafka.NewProducer(
			cfg.Kafka,
			cfg.Kafka.Topics.DonationRequest)
		if err != nil {
			diFail <- err
			return
		}

		userService := user.NewService(mongoRepo)
		donationService := donation.NewService(
			mongoRepo,
			donationProducer)

		httpServer := http.NewServer(
			cfg.Server.Http,
			userService,
			donationService)

		diDone <- &components{
			mongoRepo:        mongoRepo,
			donationProducer: donationProducer,
			userService:      userService,
			donationService:  donationService,
			httpServer:       httpServer,
		}
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("dependency injection timeout")
	case err := <-diFail:
		return nil, err
	case c := <-diDone:
		return c, nil
	}
}
