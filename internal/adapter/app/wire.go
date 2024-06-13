//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"github.com/google/wire"
	"go.uber.org/zap"
	"management/internal/adapter/config"
	psql "management/internal/adapter/storage/postgres"
	adapter_http "management/internal/adapter/transport/http"
	"management/internal/core/attribute"
	"management/internal/core/port/db"
	"management/internal/core/port/department"
	"management/internal/core/port/http"
	"management/internal/core/port/user"
	port_service "management/internal/core/service"
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
		psql.DepartmentRepositorySet,
		port_service.DepartmentServiceSet,
		psql.AttributeRepositorySet,
		port_service.AttributeServiceSet,
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
	userService user.UserServicePort,
	departmentService department.DepartmentServicePort,
	attributeService attribute.AttributeServicePort,
) (http.ServerMaker, func(), error) {
	httpServer := adapter_http.NewHTTPServer(ctx, Cfg, userService, departmentService, attributeService)
	err := httpServer.Start(ctx)
	if err != nil {
		return nil, nil, err
	}
	return httpServer, func() { httpServer.Close(ctx) }, nil
}
