package app

import (
	"context"
	"go.uber.org/zap"
	"management/internal/adapter/config"
	"management/internal/core/attribute"
	"management/internal/core/port/db"
	"management/internal/core/port/department"
	"management/internal/core/port/http"
	"management/internal/core/port/user"
	"sync"
)

type App struct {
	rw                *sync.RWMutex
	Cfg               *config.Container
	HTTP              http.ServerMaker
	PG                db.PostgresEngineMaker
	UserRepo          user.UserRepositoryPort
	UserService       user.UserServicePort
	DepartmentRepo    department.DepartmentRepositoryPort
	DepartmentService department.DepartmentServicePort
	AttributeRepo     attribute.AttributeRepositoryPort
	AttributeService  attribute.AttributeServicePort
}

func New(
	rw *sync.RWMutex,
	cfg *config.Container,
	http http.ServerMaker,
	pg db.PostgresEngineMaker,
	userRepo user.UserRepositoryPort,
	userService user.UserServicePort,
	DepartmentRepo department.DepartmentRepositoryPort,
	DepartmentService department.DepartmentServicePort,
	AttributeRepo attribute.AttributeRepositoryPort,
	AttributeService attribute.AttributeServicePort,
) *App {
	return &App{
		rw:                rw,
		Cfg:               cfg,
		HTTP:              http,
		PG:                pg,
		UserRepo:          userRepo,
		UserService:       userService,
		DepartmentRepo:    DepartmentRepo,
		DepartmentService: DepartmentService,
		AttributeRepo:     AttributeRepo,
		AttributeService:  AttributeService,
	}
}

func (a *App) Run(ctx context.Context) {
	zap.S().Info("RUNNER!")
}
