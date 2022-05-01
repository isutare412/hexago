package main

import (
	"context"
	"fmt"

	"github.com/isutare412/hexago/gateway/pkg/config"
	"github.com/isutare412/hexago/gateway/pkg/controller/http"
	"github.com/isutare412/hexago/gateway/pkg/core/service/user"
	"github.com/isutare412/hexago/gateway/pkg/infrastructure/mq"
	"github.com/isutare412/hexago/gateway/pkg/infrastructure/repo"
)

type components struct {
	mongoRepo       *repo.MongoDB
	paymentProducer *mq.KafkaProducer
	userService     *user.Service
	httpServer      *http.Server
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
		mongoRepo, err := repo.NewMongoDB(ctx, cfg.MongoDB)
		if err != nil {
			diFail <- err
			return
		}

		paymentProducer, err := mq.NewKafkaProducer(cfg.Kafka, cfg.Kafka.Topics.PaymentRequest)
		if err != nil {
			diFail <- err
			return
		}

		userService := user.NewService(mongoRepo)

		httpServer := http.NewServer(cfg.Server.Http, userService)

		diDone <- &components{
			mongoRepo:       mongoRepo,
			paymentProducer: paymentProducer,
			userService:     userService,
			httpServer:      httpServer,
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
