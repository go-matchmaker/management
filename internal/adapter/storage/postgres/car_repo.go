package psql

import (
	"context"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"management/internal/core/domain/aggregate"
	"management/internal/core/port/car"
	"management/internal/core/port/db"
)

var (
	_                car.CarRepositoryPort = (*CarRepository)(nil)
	CarRepositorySet                       = wire.NewSet(NewCarRepository)
)

type CarRepository struct {
	dbPool *pgxpool.Pool
}

func NewCarRepository(em db.PostgresEngineMaker) car.CarRepositoryPort {
	return &CarRepository{
		dbPool: em.GetDB(),
	}
}

func (cr *CarRepository) Insert(ctx context.Context, aggregateCar *aggregate.Car) (*string, error) {
	query := `INSERT INTO cars (plate, model_year, km, brand_id, model_id, color_id, fuel_id, transmission_id, situation, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	var id string
	err := cr.dbPool.QueryRow(ctx, query, aggregateCar.CarInfo.Plate, aggregateCar.CarInfo.ModelYear, aggregateCar.CarInfo.KM, aggregateCar.Brand.ID, aggregateCar.Model.ID, aggregateCar.Color.ID, aggregateCar.Fuel.ID, aggregateCar.Transmission.ID, aggregateCar.CarInfo.Situation, aggregateCar.CarInfo.CreatedAt, aggregateCar.CarInfo.UpdatedAt).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
