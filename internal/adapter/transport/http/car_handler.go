package http

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"management/internal/dto"
)

func (s *server) CreateCar(c fiber.Ctx) error {
	body := c.Body()
	if body == nil {
		return s.errorResponse(c, "Invalid request", nil, nil, 400)
	}

	carBody := new(dto.CreateCarRequest)
	err := json.Unmarshal(body, carBody)
	if err != nil {
		return s.errorResponse(c, "Invalid request", err, nil, 400)
	}

	car := converter.CarCreateDtoToModel(carBody)
	id, err := s.carService.CreateCar(c.Context(), &car)
	if err != nil {
		return s.errorResponse(c, "Failed to create car", err, nil, 400)
	}

	return s.successResponse(c, id, "Car created successfully", fiber.StatusOK)
}

func (s *server) CreateBrand(c fiber.Ctx) error {
	body := c.Body()
	if body == nil {
		return s.errorResponse(c, "Invalid request", nil, nil, 400)
	}

	brandBody := new(dto.CreateBrandRequest)
	err := json.Unmarshal(body, brandBody)
	if err != nil {
		return s.errorResponse(c, "Invalid request", err, nil, 400)
	}

	brand := converter.BrandCreateDtoToModel(brandBody)
	id, err := s.carService.CreateBrand(c.Context(), &brand)
	if err != nil {
		return s.errorResponse(c, "Failed to create brand", err, nil, 400)
	}

	return s.successResponse(c, id, "Brand created successfully", fiber.StatusOK)
}

func (s *server) CreateModel(c fiber.Ctx) error {
	body := c.Body()
	if body == nil {
		return s.errorResponse(c, "Invalid request", nil, nil, 400)
	}

	modelBody := new(dto.CreateModelRequest)
	err := json.Unmarshal(body, modelBody)
	if err != nil {
		return s.errorResponse(c, "Invalid request", err, nil, 400)
	}

	model := converter.ModelCreateDtoToModel(modelBody)
	id, err := s.carService.CreateModel(c.Context(), &model, modelBody.BrandID)
	if err != nil {
		return s.errorResponse(c, "Failed to create model", err, nil, 400)
	}

	return s.successResponse(c, id, "Model created successfully", fiber.StatusOK)
}

func (s *server) CreateColor(c fiber.Ctx) error {
	body := c.Body()
	if body == nil {
		return s.errorResponse(c, "Invalid request", nil, nil, 400)
	}

	colorBody := new(dto.CreateColorRequest)
	err := json.Unmarshal(body, colorBody)
	if err != nil {
		return s.errorResponse(c, "Invalid request", err, nil, 400)
	}

	color := converter.ColorCreateDtoToModel(colorBody)
	id, err := s.carService.CreateColor(c.Context(), &color)
	if err != nil {
		return s.errorResponse(c, "Failed to create color", err, nil, 400)
	}

	return s.successResponse(c, id, "Color created successfully", fiber.StatusOK)
}

func (s *server) CreateFuel(c fiber.Ctx) error {
	body := c.Body()
	if body == nil {
		return s.errorResponse(c, "Invalid request", nil, nil, 400)
	}

	fuelBody := new(dto.CreateFuelRequest)
	err := json.Unmarshal(body, fuelBody)
	if err != nil {
		return s.errorResponse(c, "Invalid request", err, nil, 400)
	}

	fuel := converter.FuelCreateDtoToModel(fuelBody)
	id, err := s.carService.CreateFuel(c.Context(), &fuel)
	if err != nil {
		return s.errorResponse(c, "Failed to create fuel", err, nil, 400)
	}

	return s.successResponse(c, id, "Fuel created successfully", fiber.StatusOK)
}

func (s *server) CreateTransmission(c fiber.Ctx) error {
	body := c.Body()

	if body == nil {
		return s.errorResponse(c, "Invalid request", nil, nil, 400)
	}

	transmissionBody := new(dto.CreateTransmissionRequest)
	err := json.Unmarshal(body, transmissionBody)
	if err != nil {
		return s.errorResponse(c, "Invalid request", err, nil, 400)
	}

	transmission := converter.TransmissionCreateDtoToModel(transmissionBody)
	id, err := s.carService.CreateTransmission(c.Context(), &transmission)
	if err != nil {
		return s.errorResponse(c, "Failed to create transmission", err, nil, 400)
	}

	return s.successResponse(c, id, "Transmission created successfully", fiber.StatusOK)
}
