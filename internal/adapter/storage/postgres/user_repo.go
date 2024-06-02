package psql

import (
	"context"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"management/internal/core/domain/aggregate"
	"management/internal/core/port/db"
	"management/internal/core/port/user"
)

var (
	_                 user.UserRepositoryPort = (*UserRepository)(nil)
	UserRepositorySet                         = wire.NewSet(NewUserRepository)
)

type UserRepository struct {
	dbPool *pgxpool.Pool
}

func NewUserRepository(em db.PostgresEngineMaker) user.UserRepositoryPort {
	return &UserRepository{
		dbPool: em.GetDB(),
	}
}

func (r *UserRepository) Insert(ctx context.Context, userAggregate *aggregate.UserPermission) (*string, error) {
	query := `INSERT INTO users (role, name, surname, email, phone_number, password, department_id, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	var id string
	err := r.dbPool.QueryRow(ctx, query,
		userAggregate.User.Role,
		userAggregate.User.Name,
		userAggregate.User.Surname,
		userAggregate.User.Email,
		userAggregate.User.PhoneNumber,
		userAggregate.User.Password,
		userAggregate.DepartmentName,
		userAggregate.User.CreatedAt,
		userAggregate.User.UpdatedAt).Scan(&id)

	if err != nil {
		return nil, err
	}

	userAggregate.User.ID = id
	return &id, nil
}
