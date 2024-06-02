package attribute

import (
	"context"
	"management/internal/core/domain/aggregate"
)

type AttributeRepositoryPort interface {
	Insert(ctx context.Context, userAggregate *aggregate.UserPermission) ([]string, error)
	//Update(ctx context.Context, user *entity.User) (*entity.User, error)
	//UpdatePassword(ctx context.Context, id string, password string) (*entity.User, error)
	//GetByID(ctx context.Context, id string) (entity.User, error)
	//DeleteOne(ctx context.Context, id string) error
	//DeleteAll(ctx context.Context) error
}

type AttributeServicePort interface {
	CreateAttribute(ctx context.Context, userAggregate *aggregate.UserPermission) ([]string, error)
	//DeleteOne(ctx context.Context, id string) error
	// GetByID(ctx context.Context, id string) (*model.User, error)
	// RefreshToken(ctx context.Context, userID string) (string, error)
	// ChangePassword(ctx context.Context, id string, req *dto.ChangePasswordReq) error
}
