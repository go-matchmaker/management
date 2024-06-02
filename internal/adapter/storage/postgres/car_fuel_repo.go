package psql

import (
	"context"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"management/internal/core/domain/entity"
	"management/internal/core/port/car"
	"management/internal/core/port/db"
)

var (
	_                    car.FuelRepositoryPort = (*CarFuelRepository)(nil)
	CarFuelRepositorySet                        = wire.NewSet(NewCarRepository)
)

type CarFuelRepository struct {
	dbPool *pgxpool.Pool
}

func NewCarFuelRepository(em db.PostgresEngineMaker) car.FuelRepositoryPort {
	return &CarFuelRepository{
		dbPool: em.GetDB(),
	}
}

func (cr *CarFuelRepository) Insert(ctx context.Context, fuel *entity.Fuel) (*string, error) {
	query := `INSERT INTO models (name, situation, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id string
	err := cr.dbPool.QueryRow(ctx, query, fuel.Name, fuel.Situation, fuel.CreatedAt, fuel.UpdatedAt).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
