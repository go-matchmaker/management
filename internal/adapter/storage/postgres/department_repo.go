package psql

import (
	"context"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"management/internal/core/domain/entity"
	"management/internal/core/port/db"
	"management/internal/core/port/department"
)

var (
	_                       department.DepartmentRepositoryPort = (*DepartmentRepository)(nil)
	DepartmentRepositorySet                                     = wire.NewSet(NewDepartmentRepository)
)

type DepartmentRepository struct {
	dbPool *pgxpool.Pool
}

func NewDepartmentRepository(em db.PostgresEngineMaker) department.DepartmentRepositoryPort {
	return &DepartmentRepository{
		dbPool: em.GetDB(),
	}
}

func (r *DepartmentRepository) Insert(ctx context.Context, departmentModel *entity.Department) (*string, error) {
	query := `INSERT INTO departments (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id`
	var returnedID string
	err := r.dbPool.QueryRow(ctx, query, departmentModel.Name, departmentModel.CreatedAt, departmentModel.UpdatedAt).Scan(&returnedID)
	if err != nil {
		return nil, err
	}

	departmentModel.ID = returnedID
	return &returnedID, nil
}
