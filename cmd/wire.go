//go:build wireinject

package main

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/isutare412/hexago/pkg/config"
	"github.com/isutare412/hexago/pkg/core/port"
	"github.com/isutare412/hexago/pkg/core/service/student"
	"github.com/isutare412/hexago/pkg/infrastructure/repo"
)

type beans struct {
	mongoRepo *repo.MongoDB
	stuSvc    port.StudentService
}

func dependencyInjection(
	ctx context.Context,
	cfg *config.Config,
) (*beans, error) {
	wire.Build(
		repo.NewMongoDB,
		student.NewService,
		wire.Bind(new(port.StudentRepository), new(*repo.MongoDB)),
		wire.Struct(new(beans),
			"mongoRepo",
			"stuSvc",
		),
	)
	return nil, fmt.Errorf("wire unimplemented")
}
