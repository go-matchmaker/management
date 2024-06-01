package service

import (
	"context"
	"github.com/google/wire"
	"management/internal/core/domain/entity"
	"management/internal/core/port/department"
)

var (
	_                    department.DepartmentServicePort = (*DepartmentService)(nil)
	DepartmentServiceSet                                  = wire.NewSet(NewDepartmentService)
)

type DepartmentService struct {
	departmentRepo department.DepartmentRepositoryPort
}

func NewDepartmentService(departmentRepo department.DepartmentRepositoryPort) department.DepartmentServicePort {
	return &DepartmentService{
		departmentRepo,
	}
}

func (us *DepartmentService) CreateDepartment(ctx context.Context, departmentModel *entity.Department) (*string, error) {
	id, err := us.departmentRepo.Insert(ctx, departmentModel)
	if err != nil {
		return nil, err
	}

	return id, nil
}
