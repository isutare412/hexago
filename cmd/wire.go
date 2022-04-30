package main

import (
	"context"

	"github.com/isutare412/hexago/pkg/config"
	"github.com/isutare412/hexago/pkg/core/port"
	"github.com/isutare412/hexago/pkg/core/service/student"
	"github.com/isutare412/hexago/pkg/infrastructure/repo"
)

type beans struct {
	mongoRepo      *repo.MongoDB
	studentService port.StudentService
}

func dependencyInjection(
	ctx context.Context,
	cfg *config.Config,
) (*beans, error) {
	mongoRepo, err := repo.NewMongoDB(ctx, cfg.MongoDB)
	if err != nil {
		return nil, err
	}

	studentService := student.NewService(mongoRepo)

	return &beans{
		mongoRepo:      mongoRepo,
		studentService: studentService,
	}, nil
}
