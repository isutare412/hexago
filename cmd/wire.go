package main

import (
	"context"

	"github.com/isutare412/hexago/pkg/config"
	"github.com/isutare412/hexago/pkg/core/service/user"
	"github.com/isutare412/hexago/pkg/infrastructure/repo"
)

type beans struct {
	mongoRepo   *repo.MongoDB
	userService *user.Service
}

func dependencyInjection(
	ctx context.Context,
	cfg *config.Config,
) (*beans, error) {
	mongoRepo, err := repo.NewMongoDB(ctx, cfg.MongoDB)
	if err != nil {
		return nil, err
	}

	userService := user.NewService(mongoRepo)

	return &beans{
		mongoRepo:   mongoRepo,
		userService: userService,
	}, nil
}
