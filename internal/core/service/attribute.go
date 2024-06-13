package service

import (
	"context"
	"github.com/google/wire"
	"management/internal/core/attribute"
	"management/internal/core/domain/aggregate"
)

var (
	_                   attribute.AttributeServicePort = (*AttributeService)(nil)
	AttributeServiceSet                                = wire.NewSet(NewAttributeService)
)

type AttributeService struct {
	attributeRepo attribute.AttributeRepositoryPort
}

func NewAttributeService(attributeRepo attribute.AttributeRepositoryPort) attribute.AttributeServicePort {
	return &AttributeService{
		attributeRepo,
	}
}

func (us *AttributeService) CreateAttribute(ctx context.Context, userAggregate *aggregate.UserPermission) ([]string, error) {
	ids, err := us.attributeRepo.Insert(ctx, userAggregate)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
