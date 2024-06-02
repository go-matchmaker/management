package converter

import (
	"management/internal/core/domain/entity"
	"management/internal/core/domain/valueobject"
	"management/internal/dto"
	"time"
)

func UserCreateDtoToModel(dtoData *dto.CreateUserRequest) (entity.User, map[string]valueobject.Attribute) {
	user := entity.User{
		Role:        entity.Role(dtoData.User.Role),
		Name:        dtoData.User.Name,
		Surname:     dtoData.User.Surname,
		Email:       dtoData.User.Email,
		PhoneNumber: dtoData.User.PhoneNumber,
		Password:    dtoData.User.Password,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	permissions := dtoData.Permissions
	attributes := make(map[string]valueobject.Attribute)
	for key, value := range permissions {
		attributes[key] = valueobject.Attribute(value)
	}

	return user, attributes
}

func DepartmentCreateDtoToModel(dtoData *dto.CreateDepartmentRequest) entity.Department {
	return entity.Department{
		Name:      dtoData.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
