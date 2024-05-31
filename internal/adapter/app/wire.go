//go:build wireinject
// +build wireinject

package app

import (
	"management/internal/adapter/management/paseto"
	"management/internal/adapter/config"
	psql "management/internal/adapter/storage/postgres"
	adapter_http "management/internal/adapter/transport/http"
	"management/internal/core/port/management"
	"management/internal/core/port/db"
	"management/internal/core/port/http"
	"management/internal/core/port/user"
	port_service "management/internal/core/service"
	"context"
	"github.com/google/wire"
	"go.uber.org/zap"
	"sync"
)

func InitApp(
	ctx context.Context,
	wg *sync.WaitGroup,
	rw *sync.RWMutex,
	Cfg *config.Container,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		httpServerFunc,
		psql.UserRepositorySet,
		port_service.UserServiceSet,
		paseto.PasetoSet,
	))
}

func dbEngineFunc(
	ctx context.Context,
	Cfg *config.Container) (db.PostgresEngineMaker, func(), error) {
	psqlDb := psql.NewDB(Cfg)
	err := psqlDb.Start(ctx)
	if err != nil {
		zap.S().Fatal("failed to start db:", err)
	}

	if err != nil {
		zap.S().Fatal("failed to migrate db:", err)
	}
	return psqlDb, func() { psqlDb.Close(ctx) }, nil
}

func httpServerFunc(
	ctx context.Context,
	Cfg *config.Container,
	UserService user.UserServicePort,
	tokenMaker management.TokenMaker,
) (http.ServerMaker, func(), error) {
	httpServer := adapter_http.NewHTTPServer(ctx, Cfg, UserService, tokenMaker)
	err := httpServer.Start(ctx)
	if err != nil {
		return nil, nil, err
	}
	return httpServer, func() { httpServer.Close(ctx) }, nil
}
