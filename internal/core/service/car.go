package service

import (
	"context"
	"github.com/google/wire"
	"management/internal/core/domain/aggregate"
	"management/internal/core/domain/entity"
	"management/internal/core/port/car"
)

var (
	_             car.CarServicePort = (*CarService)(nil)
	CarServiceSet                    = wire.NewSet(NewCarService)
)

type CarService struct {
	carRepo          car.CarRepositoryPort
	brandRepo        car.BrandRepositoryPort
	modelRepo        car.ModelRepositoryPort
	transmissionRepo car.TransmissionRepositoryPort
	colorRepo        car.ColorRepositoryPort
	fuelRepo         car.FuelRepositoryPort
}

func NewCarService(carRepo car.CarRepositoryPort, brandRepo car.BrandRepositoryPort, modelRepo car.ModelRepositoryPort, transmissionRepo car.TransmissionRepositoryPort, colorRepo car.ColorRepositoryPort, fuelRepo car.FuelRepositoryPort) car.CarServicePort {
	return &CarService{
		carRepo,
		brandRepo,
		modelRepo,
		transmissionRepo,
		colorRepo,
		fuelRepo,
	}
}

func (us *CarService) CreateCar(ctx context.Context, carAggragate *aggregate.Car) (*string, error) {
	id, err := us.carRepo.Insert(ctx, carAggragate)
	if err != nil {
		return nil, err
	}

	return id, nil

}

func (us *CarService) CreateBrand(ctx context.Context, carBrand *entity.Brand) (*string, error) {
	id, err := us.brandRepo.Insert(ctx, carBrand)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (us *CarService) CreateModel(ctx context.Context, carModel *entity.Model, brandID string) (*string, error) {
	id, err := us.modelRepo.Insert(ctx, carModel, brandID)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (us *CarService) CreateColor(ctx context.Context, carColor *entity.Color) (*string, error) {
	id, err := us.colorRepo.Insert(ctx, carColor)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (us *CarService) CreateFuel(ctx context.Context, carFuel *entity.Fuel) (*string, error) {
	id, err := us.fuelRepo.Insert(ctx, carFuel)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (us *CarService) CreateTransmission(ctx context.Context, carTransmission *entity.Transmission) (*string, error) {
	id, err := us.transmissionRepo.Insert(ctx, carTransmission)
	if err != nil {
		return nil, err
	}

	return id, nil
}
