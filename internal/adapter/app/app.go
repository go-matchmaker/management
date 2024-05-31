package app

import (
	"context"
	"go.uber.org/zap"
	"management/internal/adapter/config"
	"management/internal/core/port/db"
	"management/internal/core/port/http"
	"management/internal/core/port/management"
	"management/internal/core/port/user"
	"sync"
)

type App struct {
	rw          *sync.RWMutex
	Cfg         *config.Container
	HTTP        http.ServerMaker
	Token       management.TokenMaker
	PG          db.PostgresEngineMaker
	UserRepo    user.UserRepositoryPort
	UserService user.UserServicePort
}

func New(
	rw *sync.RWMutex,
	cfg *config.Container,
	http http.ServerMaker,
	token management.TokenMaker,
	pg db.PostgresEngineMaker,
	userRepo user.UserRepositoryPort,
	userService user.UserServicePort,
) *App {
	return &App{
		rw:          rw,
		Cfg:         cfg,
		HTTP:        http,
		Token:       token,
		PG:          pg,
		UserRepo:    userRepo,
		UserService: userService,
	}
}

func (a *App) Run(ctx context.Context) {
	zap.S().Info("RUNNER!")
}