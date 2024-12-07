//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	_ "github.com/lib/pq"
	http "product-service/internal/api"
	"product-service/internal/api/handler"
	"product-service/internal/config"
	"product-service/internal/db"
	"product-service/internal/repository"
	"product-service/internal/service"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		handler.NewProductHandler,
		repository.NewProductRepository,
		service.NewProductService,
		http.NewServer,
	)
	return &http.Server{}, nil
}
