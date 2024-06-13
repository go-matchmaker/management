package aggregate

import (
	"management/internal/core/domain/entity"
	"management/internal/core/domain/valueobject"
)

type UserPermission struct {
	User           entity.User                      `json:"user"`
	DepartmentName string                           `json:"department_name"`
	Attributes     map[string]valueobject.Attribute `json:"permissions"`
}

func NewUserPermission(user entity.User, attributes map[string]valueobject.Attribute, departmentName string) *UserPermission {
	return &UserPermission{
		User:           user,
		DepartmentName: departmentName,
		Attributes:     attributes,
	}
}
