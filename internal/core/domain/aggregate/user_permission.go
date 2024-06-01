package aggregate

import (
	"management/internal/core/domain/entity"
	"management/internal/core/domain/valueobject"
)

type UserPermission struct {
	User         entity.User                      `json:"user"`
	DepartmentID string                           `json:"department"`
	Attributes   map[string]valueobject.Attribute `json:"permissions"`
}

func NewUserPermission(user entity.User, attributes map[string]valueobject.Attribute, departmentID string) *UserPermission {
	return &UserPermission{
		User:         user,
		DepartmentID: departmentID,
		Attributes:   attributes,
	}
}
