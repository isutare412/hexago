package main

import (
	"context"

	"github.com/isutare412/hexago/gateway/pkg/config"
	"github.com/isutare412/hexago/gateway/pkg/controller/http"
	"github.com/isutare412/hexago/gateway/pkg/core/service/user"
	"github.com/isutare412/hexago/gateway/pkg/infrastructure/repo"
)

type components struct {
	mongoRepo   *repo.MongoDB
	userService *user.Service
	httpServer  *http.Server
}

func dependencyInjection(
	ctx context.Context,
	cfg *config.Config,
) (*components, error) {
	mongoRepo, err := repo.NewMongoDB(ctx, cfg.MongoDB)
	if err != nil {
		return nil, err
	}

	userService := user.NewService(mongoRepo)

	httpServer := http.NewServer(cfg.Server.Http, userService)

	return &components{
		mongoRepo:   mongoRepo,
		userService: userService,
		httpServer:  httpServer,
	}, nil
}
