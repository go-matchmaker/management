package http

func (s *server) SetupRouter() {
	s.hrManagementRouter()
}

func (s *server) hrManagementRouter() {
	route := s.app.Group("/hr-management")
	route.Post("/create-user", s.CreateUser)
	route.Post("/create-department", s.CreateDepartment)
}

func (s *server) carManagementRouter() {
	route := s.app.Group("/car-management")
	route.Post("/create-car", s.CreateCar)
	route.Post("/create-brand", s.CreateBrand)
	route.Post("/create-model", s.CreateModel)
	route.Post("/create-color", s.CreateColor)
	route.Post("/create-fuel", s.CreateFuel)
	route.Post("/create-transmission", s.CreateTransmission)
}
