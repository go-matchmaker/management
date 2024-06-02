package http

import "time"

func (s *server) SetupRouter() {
	s.managementRouter()
}

func (s *server) managementRouter() {
	route := s.app.Group("/management")
	route.Post("/create-user", s.CreateUser, s.RateLimiter(5, time.Minute))
	route.Post("/create-department", s.CreateDepartment, s.RateLimiter(5, time.Minute))
}
