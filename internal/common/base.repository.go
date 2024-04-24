package common

import "gorm.io/gorm"

// Repository defines the interface for a generic repository
type Repository interface {
	Create(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	FindAll(models interface{}) error
	FindByID(model interface{}, id string) error
}

// BaseRepository is a generic implementation of the Repository interface
type BaseRepository struct {
	db *gorm.DB
}

// NewBaseRepository creates a new instance of BaseRepository
func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

// Create adds a new record to the database
func (r *BaseRepository) Create(model interface{}) error {
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	return nil
}

// Update updates an existing record in the database
func (r *BaseRepository) Update(model interface{}) error {
	if err := r.db.Save(model).Error; err != nil {
		return err
	}
	return nil
}

// Delete removes a record from the database
func (r *BaseRepository) Delete(model interface{}) error {
	if err := r.db.Delete(model).Error; err != nil {
		return err
	}
	return nil
}

// FindAll retrieves all records of a model from the database
func (r *BaseRepository) FindAll(models interface{}) error {
	if err := r.db.Find(models).Error; err != nil {
		return err
	}
	return nil
}

// FindByID retrieves a record by its ID from the database
func (r *BaseRepository) FindByID(model interface{}, id uint) error {
	if err := r.db.First(model, id).Error; err != nil {
		return err
	}
	return nil
}
