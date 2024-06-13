package service

import (
	"context"
	"github.com/google/wire"
	"management/internal/core/domain/aggregate"
	"management/internal/core/port/user"
)

var (
	_              user.UserServicePort = (*UserService)(nil)
	UserServiceSet                      = wire.NewSet(NewUserService)
)

type UserService struct {
	userRepo user.UserRepositoryPort
}

func NewUserService(userRepo user.UserRepositoryPort) user.UserServicePort {
	return &UserService{
		userRepo,
	}
}

func (us *UserService) CreateUser(ctx context.Context, userAggregate *aggregate.UserPermission) (*string, error) {
	id, err := us.userRepo.Insert(ctx, userAggregate)
	if err != nil {
		return nil, err
	}

	return id, nil
}
