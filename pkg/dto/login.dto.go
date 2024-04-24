package dto

// dto for login
type LoginDto struct {
	Email    string `json:"email" validate:"required,email,min=1"`
	Password string `json:"password" validate:"required,min=6"`
}
