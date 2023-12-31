//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/abhinandpn/StockX/pkg/api"
	handler "github.com/abhinandpn/StockX/pkg/api/handler"
	config "github.com/abhinandpn/StockX/pkg/config"
	db "github.com/abhinandpn/StockX/pkg/db"
	repository "github.com/abhinandpn/StockX/pkg/repository"
	usecase "github.com/abhinandpn/StockX/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
