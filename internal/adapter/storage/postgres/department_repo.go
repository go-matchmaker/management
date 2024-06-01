package psql

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"management/internal/core/domain/entity"
	"management/internal/core/port/db"
	"management/internal/core/port/department"
	"time"
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

func (r *DepartmentRepository) getNextDepartmentIndex(ctx context.Context) (int, error) {
	var userIndex int
	err := r.dbPool.QueryRow(ctx, "SELECT COUNT(*) FROM departments").Scan(&userIndex)
	if err != nil {
		return 0, err
	}
	return userIndex + 1, nil
}

func (r *DepartmentRepository) Insert(ctx context.Context, departmentModel *entity.Department) (*string, error) {
	departmentIndex, err := r.getNextDepartmentIndex(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	formattedDate := now.Format("20060102")
	id := fmt.Sprintf("%s%d", formattedDate, departmentIndex)
	departmentModel.ID = id
	query := `INSERT INTO departments (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

	var returnedID string
	err = r.dbPool.QueryRow(ctx, query, departmentModel.ID, departmentModel.Name, departmentModel.CreatedAt, departmentModel.UpdatedAt).Scan(&returnedID)
	if err != nil {
		return nil, err
	}

	return &returnedID, nil
}
