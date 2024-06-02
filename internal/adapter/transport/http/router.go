package http

import "time"

func (s *server) SetupRouter() {
	s.hrManagementRouter()
}

func (s *server) hrManagementRouter() {
	route := s.app.Group("/hr-management")
	route.Post("/create-user", s.CreateUser, s.RateLimiter(5, time.Minute))
	route.Post("/create-department", s.CreateDepartment, s.RateLimiter(5, time.Minute))
}

func (s *server) carManagementRouter() {
	route := s.app.Group("/car-management")
	route.Post("/create-car", s.CreateCar, s.RateLimiter(5, time.Minute))
	route.Post("/create-brand", s.CreateBrand, s.RateLimiter(5, time.Minute))
	route.Post("/create-model", s.CreateModel, s.RateLimiter(5, time.Minute))
	route.Post("/create-color", s.CreateColor, s.RateLimiter(5, time.Minute))
	route.Post("/create-fuel", s.CreateFuel, s.RateLimiter(5, time.Minute))
	route.Post("/create-transmission", s.CreateTransmission, s.RateLimiter(5, time.Minute))
}
