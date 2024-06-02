package psql

import (
	"context"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"management/internal/core/attribute"
	"management/internal/core/domain/aggregate"
	"management/internal/core/port/db"
)

var (
	_                      attribute.AttributeRepositoryPort = (*AttributeRepository)(nil)
	AttributeRepositorySet                                   = wire.NewSet(NewAttributeRepository)
)

type AttributeRepository struct {
	dbPool *pgxpool.Pool
}

func NewAttributeRepository(em db.PostgresEngineMaker) attribute.AttributeRepositoryPort {
	return &AttributeRepository{
		dbPool: em.GetDB(),
	}
}

func (r *AttributeRepository) Insert(ctx context.Context, aggregateModel *aggregate.UserPermission) ([]string, error) {
	attribureIds := make([]string, 0)
	for k, v := range aggregateModel.Attributes {
		query := `INSERT INTO attributes (user_id, attribute, view, search, detail, add, update, delete, export, import, can_see_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
		var id string
		err := r.dbPool.QueryRow(ctx, query,
			aggregateModel.User.ID,
			k,
			v.View,
			v.Search,
			v.Detail,
			v.Add,
			v.Update,
			v.Delete,
			v.Export,
			v.Import,
			v.CanSeePrice).Scan(&id)

		if err != nil {
			return nil, err
		}

		attribureIds = append(attribureIds, id)
	}
	return attribureIds, nil
}
