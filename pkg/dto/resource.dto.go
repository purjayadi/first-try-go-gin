package dto

import "github.com/google/uuid"

// make dto for resource
type ResourceDto struct {
	Name string `json:"name" validate:"required,min=1"`
}

type UpdateResourceDto struct {
	ID   uuid.UUID `json:"id" validate:"required,uuid"`
	Name string    `json:"name" validate:"required,min=1"`
}

type SearchResourceDto struct {
	Page     *int    `json:"page" form:"page"`
	PageSize *int    `json:"pageSize" form:"pageSize"`
	Search   *string `json:"search" form:"search"`
}
