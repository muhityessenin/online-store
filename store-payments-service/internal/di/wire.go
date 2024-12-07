//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	_ "github.com/lib/pq"
	http "payment-service/internal/api"
	"payment-service/internal/api/handler"
	"payment-service/internal/config"
	"payment-service/internal/db"
	"payment-service/internal/repository"
	"payment-service/internal/service"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		handler.NewPaymentHandler,
		repository.NewPaymentRepository,
		service.NewPaymentService,
		http.NewServer,
	)
	return &http.Server{}, nil
}
