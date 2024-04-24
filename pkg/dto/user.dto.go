package dto

type CreateUserDto struct {
	Name     string `json:"name" validate:"required,min=1"`
	Email    string `json:"email" validate:"required,email,min=1"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserDto struct {
	Name     string `json:"name" validate:"required,min=1"`
	Email    string `json:"email" validate:"required,email,min=1"`
	Password string `json:"password" validate:"required,min=6"`
}

type GetUserDto struct {
	PaginationDto
}
