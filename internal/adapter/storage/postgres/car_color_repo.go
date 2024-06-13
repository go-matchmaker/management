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
	_                     car.ColorRepositoryPort = (*CarColorRepository)(nil)
	CarColorRepositorySet                         = wire.NewSet(NewCarRepository)
)

type CarColorRepository struct {
	dbPool *pgxpool.Pool
}

func NewCarColorRepository(em db.PostgresEngineMaker) car.ColorRepositoryPort {
	return &CarColorRepository{
		dbPool: em.GetDB(),
	}
}

func (cr *CarColorRepository) Insert(ctx context.Context, color *entity.Color) (*string, error) {
	query := `INSERT INTO models (name, color_code, situation, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id string
	err := cr.dbPool.QueryRow(ctx, query, color.Name, color.ColorCode, color.Situation, color.CreatedAt, color.UpdatedAt).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
