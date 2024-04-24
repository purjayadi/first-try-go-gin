package dto

import "github.com/google/uuid"

// make dto for package
type PackageDto struct {
	Name        string  `json:"name" validate:"required,min=1"`
	Description string  `json:"description" validate:"required,min=1"`
	Image       string  `json:"image" validate:"required,min=1"`
	Price       float64 `json:"price" validate:"required,min=1"`
}

type UpdatePackageDto struct {
	PackageDto
	ID uuid.UUID `json:"id"`
}

type GetPackageDto struct {
	PaginationDto
	Search *string `json:"search" form:"search"`
}
