package main

import (
	"context"

	"github.com/isutare412/hexago/payment/pkg/config"
)

type Components struct {
}

func dependencyInjection(
	ctx context.Context,
	cfg *config.Config,
) (*Components, error) {
	return &Components{}, nil
}
