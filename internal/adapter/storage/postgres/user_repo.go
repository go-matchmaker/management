package psql

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"management/internal/core/domain/aggregate"
	"management/internal/core/port/db"
	"management/internal/core/port/user"
	"time"
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

func (r *UserRepository) getNextUserIndex(ctx context.Context) (int, error) {
	var userIndex int
	err := r.dbPool.QueryRow(ctx, "SELECT COUNT(*) FROM users").Scan(&userIndex)
	if err != nil {
		return 0, err
	}
	return userIndex + 1, nil
}

func (r *UserRepository) Insert(ctx context.Context, userAggregate *aggregate.UserPermission) (*string, error) {
	userIndex, err := r.getNextUserIndex(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	formattedDate := now.Format("20060102")

	id := fmt.Sprintf("%s%d", formattedDate, userIndex)
	query := `INSERT INTO users (id, role, name, surname, email, phone_number, password, department_id created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, role, name, surname, email, phone_number, password_hash, created_at, updated_at`
	queryRow := r.dbPool.QueryRow(ctx, query, userAggregate.User.ID, userAggregate.User.Role, userAggregate.User.Name, userAggregate.User.Surname, userAggregate.User.Email, userAggregate.User.PhoneNumber, userAggregate.User.Password, userAggregate.DepartmentID, userAggregate.User.CreatedAt, userAggregate.User.UpdatedAt)
	err = queryRow.Scan(&userAggregate.User.ID, &userAggregate.User.Role, &userAggregate.User.Name, &userAggregate.User.Surname, &userAggregate.User.Email, &userAggregate.User.PhoneNumber, &userAggregate.User.Password, &userAggregate.User.CreatedAt, &userAggregate.User.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
