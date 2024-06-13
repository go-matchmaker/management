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
	_                     car.ModelRepositoryPort = (*CarModelRepository)(nil)
	CarModelRepositorySet                         = wire.NewSet(NewCarRepository)
)

type CarModelRepository struct {
	dbPool *pgxpool.Pool
}

func NewCarModelRepository(em db.PostgresEngineMaker) car.ModelRepositoryPort {
	return &CarModelRepository{
		dbPool: em.GetDB(),
	}
}

func (cr *CarModelRepository) Insert(ctx context.Context, model *entity.Model, brandID string) (*string, error) {
	query := `INSERT INTO models (brand_id, name, situation, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id string
	err := cr.dbPool.QueryRow(ctx, query, brandID, model.Name, model.Situation, model.CreatedAt, model.UpdatedAt).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
