package http

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"management/internal/converter"
	"management/internal/core/domain/aggregate"
	"management/internal/dto"
)

func (s *server) CreateUser(c fiber.Ctx) error {
	body := c.Body()
	if body == nil {
		return s.errorResponse(c, "Invalid request", nil, nil, 400)
	}

	userBody := new(dto.CreateUserRequest)
	err := json.Unmarshal(body, userBody)
	if err != nil {
		return s.errorResponse(c, "Invalid request", err, nil, 400)
	}

	user, permissions := converter.UserCreateDtoToModel(userBody)
	userAggregate := aggregate.NewUserPermission(user, permissions, userBody.DepartmentID)

	id, err := s.userService.CreateUser(c.Context(), userAggregate)
	if err != nil {
		return s.errorResponse(c, "Failed to create user", err, nil, 400)
	}

	_, err = s.attributeService.CreateAttribute(c.Context(), userAggregate)
	if err != nil {
		return s.errorResponse(c, "Failed to create attribute", err, nil, 400)
	}

	return s.successResponse(c, id, "User created successfully", fiber.StatusOK)
}

func (s *server) CreateDepartment(c fiber.Ctx) error {
	body := c.Body()
	if body == nil {
		return s.errorResponse(c, "Invalid request", nil, nil, 400)
	}

	departmentBody := new(dto.CreateDepartmentRequest)
	err := json.Unmarshal(body, departmentBody)
	if err != nil {
		return s.errorResponse(c, "Invalid request", err, nil, 400)
	}

	department := converter.DepartmentCreateDtoToModel(departmentBody)
	id, err := s.departmentService.CreateDepartment(c.Context(), &department)
	if err != nil {
		return s.errorResponse(c, "Failed to create department", err, nil, 400)
	}

	return s.successResponse(c, id, "Department created successfully", fiber.StatusOK)
}
