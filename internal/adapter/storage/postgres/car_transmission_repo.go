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
	_                            car.TransmissionRepositoryPort = (*CarTransmissionRepository)(nil)
	CarTransmissionRepositorySet                                = wire.NewSet(NewCarRepository)
)

type CarTransmissionRepository struct {
	dbPool *pgxpool.Pool
}

func NewCarTransmissionRepository(em db.PostgresEngineMaker) car.TransmissionRepositoryPort {
	return &CarTransmissionRepository{
		dbPool: em.GetDB(),
	}
}

func (cr *CarTransmissionRepository) Insert(ctx context.Context, color *entity.Transmission) (*string, error) {
	query := `INSERT INTO models (name, situation, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id string
	err := cr.dbPool.QueryRow(ctx, query, color.Name, color.Situation, color.CreatedAt, color.UpdatedAt).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
