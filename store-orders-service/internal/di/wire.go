//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	_ "github.com/lib/pq"
	http "order-service/internal/api"
	"order-service/internal/api/handler"
	"order-service/internal/config"
	"order-service/internal/db"
	"order-service/internal/repository"
	"order-service/internal/service"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		handler.NewOrderHandler,
		repository.NewOrderRepository,
		service.NewOrderService,
		http.NewServer,
	)
	return &http.Server{}, nil
}
