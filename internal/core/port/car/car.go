package car

import (
	"context"
	"management/internal/core/domain/aggregate"
	"management/internal/core/domain/entity"
)

type CarRepositoryPort interface {
	Insert(ctx context.Context, carAggragate *aggregate.Car) (*string, error)
	//Update(ctx context.Context, user *entity.User) (*entity.User, error)
	//UpdatePassword(ctx context.Context, id string, password string) (*entity.User, error)
	//GetByID(ctx context.Context, id string) (entity.User, error)
	//DeleteOne(ctx context.Context, id string) error
	//DeleteAll(ctx context.Context) error
}

type BrandRepositoryPort interface {
	Insert(ctx context.Context, carBrand *entity.Brand) (*string, error)
	//Update(ctx context.Context, user **entity.User) (**entity.User, error)
	//UpdatePassword(ctx context.Context, id string, password string) (**entity.User, error)
	//GetByID(ctx context.Context, id string) (*entity.User, error)
	//DeleteOne(ctx context.Context, id string) error
	//DeleteAll(ctx context.Context) error
}

type ModelRepositoryPort interface {
	Insert(ctx context.Context, carBrand *entity.Model, brandID string) (*string, error)
	//Update(ctx context.Context, user **entity.User) (**entity.User, error)
	//UpdatePassword(ctx context.Context, id string, password string) (**entity.User, error)
	//GetByID(ctx context.Context, id string) (*entity.User, error)
	//DeleteOne(ctx context.Context, id string) error
	//DeleteAll(ctx context.Context) error
}

type ColorRepositoryPort interface {
	Insert(ctx context.Context, carBrand *entity.Color) (*string, error)
}

type FuelRepositoryPort interface {
	Insert(ctx context.Context, carBrand *entity.Fuel) (*string, error)
}

type TransmissionRepositoryPort interface {
	Insert(ctx context.Context, carBrand *entity.Transmission) (*string, error)
}

type CarServicePort interface {
	CreateCar(ctx context.Context, carAggragate *aggregate.Car) (*string, error)
	CreateBrand(ctx context.Context, carBrand *entity.Brand) (*string, error)
	CreateModel(ctx context.Context, carBrand *entity.Model, brandID string) (*string, error)
	CreateColor(ctx context.Context, carBrand *entity.Color) (*string, error)
	CreateFuel(ctx context.Context, carBrand *entity.Fuel) (*string, error)
	CreateTransmission(ctx context.Context, carBrand *entity.Transmission) (*string, error)
}
