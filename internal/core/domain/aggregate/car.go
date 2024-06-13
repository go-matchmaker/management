package aggregate

import "management/internal/core/domain/entity"

type Car struct {
	CarInfo      entity.CarInfo      `json:"car"`
	Brand        entity.Brand        `json:"brand"`
	Model        entity.Model        `json:"model"`
	Color        entity.Color        `json:"color"`
	Fuel         entity.Fuel         `json:"fuel"`
	Transmission entity.Transmission `json:"transmission"`
}

func NewCar(carInfo entity.CarInfo, brand entity.Brand, model entity.Model, color entity.Color, fuel entity.Fuel, transmission entity.Transmission) *Car {
	return &Car{
		CarInfo:      carInfo,
		Brand:        brand,
		Model:        model,
		Color:        color,
		Fuel:         fuel,
		Transmission: transmission,
	}
}
