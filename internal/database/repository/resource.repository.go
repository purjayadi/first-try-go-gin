package repository

import (
	"learn-go/internal/database/model"
	"learn-go/pkg/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResourceRepository interface {
	FindAll(inputUser *dto.SearchResourceDto) ([]model.Resource, error)
	FindOne(id uuid.UUID) (model.Resource, error)
	Create(InputResource *model.Resource) error
	Update(InputResource *model.Resource) error
	Delete(InputResource *model.Resource) error
}

type ResourceRepositoryImpl struct {
	db *gorm.DB
}

// find all resource with page, pageSize and search query by name
func (r *ResourceRepositoryImpl) FindAll(inputUser *dto.SearchResourceDto) ([]model.Resource, error) {
	var resources []model.Resource
	query := r.db
	if inputUser.Page != nil {
		query = query.Limit(*inputUser.Page)
	}
	if inputUser.PageSize != nil {
		query = query.Offset(*inputUser.PageSize)
	}
	if inputUser.Search != nil {
		query = query.Where("name ILIKE ?", "%"+*inputUser.Search+"%")
	}
	err := query.Find(&resources).Error
	if err != nil {
		return []model.Resource{}, err
	}
	return resources, nil
}

// find by id
func (r *ResourceRepositoryImpl) FindOne(id uuid.UUID) (model.Resource, error) {
	var resources model.Resource
	err := r.db.Where("id =?", id).First(&resources).Error
	if err != nil {
		return model.Resource{}, err
	}
	return resources, nil
}

// create new Resource
func (r *ResourceRepositoryImpl) Create(resources *model.Resource) error {
	return r.db.Create(resources).Error
}

// update Resource
func (r *ResourceRepositoryImpl) Update(resources *model.Resource) error {
	return r.db.Save(resources).Error
}

// delete Resource
func (r *ResourceRepositoryImpl) Delete(resources *model.Resource) error {
	return r.db.Delete(resources).Error
}

func ResourceRepositoryInit(db *gorm.DB) *ResourceRepositoryImpl {
	return &ResourceRepositoryImpl{
		db: db,
	}
}
