//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	_ "github.com/lib/pq"
	http "todo-list/internal/api"
	"users-service/internal/api/handler"
	"users-service/internal/config"
	"users-service/internal/db"
	"users-service/internal/repository"
	"users-service/internal/service"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		handler.NewUserHandler,
		repository.NewUserRepository,
		service.NewUserService,
		http.NewServer,
	)
	return &http.Server{}, nil
}
