package dto

// Requests
type CreateUserRequest struct {
	User         User                  `json:"user"`
	Permissions  map[string]Permission `json:"permissions"`
	DepartmentID string                `json:"department_id"`
}

// Responses
