package dto

type User struct {
	Role        string `json:"role"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type Permission struct {
	View        bool `json:"view"`
	Search      bool `json:"search"`
	Detail      bool `json:"detail"`
	Add         bool `json:"add"`
	Update      bool `json:"update"`
	Delete      bool `json:"delete"`
	Export      bool `json:"export"`
	Import      bool `json:"import"`
	CanSeePrice bool `json:"can_see_price"`
}
