//go:build wireinject
// +build wireinject

package di

import (
	http "api-gateway-service/internal/api"
	"api-gateway-service/internal/api/handler"
	"api-gateway-service/internal/config"
	"api-gateway-service/internal/db"
	"api-gateway-service/internal/repository"
	"api-gateway-service/internal/service"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		handler.NewTaskHandler,
		repository.NewTaskRepository,
		service.NewTaskService,
		http.NewServer,
	)
	return &http.Server{}, nil
}
