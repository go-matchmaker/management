package entity

import (
	"time"
)

type CarInfo struct {
	ID        string    `json:"id"`
	Plate     string    `json:"plate"`
	ModelYear int       `json:"model_year"`
	KM        int       `json:"km"`
	Situation bool      `json:"situation"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Brand struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Situation bool      `json:"situation"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Model struct {
	ID        string    `json:"id"`
	BrandID   string    `json:"brand_id"`
	Name      string    `json:"name"`
	Situation bool      `json:"situation"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Color struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	ColorCode string    `json:"color_code"`
	Situation bool      `json:"situation"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Fuel struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Situation bool      `json:"situation"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Transmission struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Situation bool      `json:"situation"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
