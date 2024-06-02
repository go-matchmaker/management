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
	_                     car.BrandRepositoryPort = (*CarBrandRepository)(nil)
	CarBrandRepositorySet                         = wire.NewSet(NewCarRepository)
)

type CarBrandRepository struct {
	dbPool *pgxpool.Pool
}

func NewCarBrandRepository(em db.PostgresEngineMaker) car.BrandRepositoryPort {
	return &CarBrandRepository{
		dbPool: em.GetDB(),
	}
}

func (cr *CarBrandRepository) Insert(ctx context.Context, brand *entity.Brand) (*string, error) {
	query := `INSERT INTO brands (name, situation, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	var id string
	err := cr.dbPool.QueryRow(ctx, query, brand.Name, brand.Situation, brand.CreatedAt, brand.UpdatedAt).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
