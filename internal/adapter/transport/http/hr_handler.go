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
	return s.successResponse(c, id, "User created successfully", fiber.StatusOK)
}
