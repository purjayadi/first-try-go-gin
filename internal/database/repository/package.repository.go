package repository

import (
	"learn-go/internal/database/model"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PackageRepository interface {
	FindAll(args ...int) ([]model.Package, error)
	FindOne(id uuid.UUID) (model.Package, error)
	Create(InputPackage *model.Package) error
	Update(InputPackage *model.Package) error
	Delete(InputPackage *model.Package) error
}

type PackageRepositoryImpl struct {
	db *gorm.DB
}

func (r *PackageRepositoryImpl) FindAll(args ...int) ([]model.Package, error) {
	var packages []model.Package
	log.Debug(args)
	// Check if limit is provided
	query := r.db
	if len(args) > 0 && args[1] > 0 {
		query = query.Limit(args[1])
	}

	// Check if offset is provided
	if len(args) > 1 && args[0] >= 0 {
		query = query.Offset(args[0])
	}
	err := query.Find(&packages).Error
	// count,
	if err != nil {
		return nil, err
	}
	return packages, nil
}

// find by id
func (r *PackageRepositoryImpl) FindOne(id uuid.UUID) (model.Package, error) {
	var Package model.Package
	err := r.db.Where("id =?", id).First(&Package).Error
	if err != nil {
		return model.Package{}, err
	}
	return Package, nil
}

// create new package
func (r *PackageRepositoryImpl) Create(Package *model.Package) error {
	return r.db.Create(Package).Error
}

// update package
func (r *PackageRepositoryImpl) Update(Package *model.Package) error {
	return r.db.Save(Package).Error
}

// delete package
func (r *PackageRepositoryImpl) Delete(Package *model.Package) error {
	return r.db.Delete(Package).Error
}

func PackageRepositoryInit(db *gorm.DB) *PackageRepositoryImpl {
	return &PackageRepositoryImpl{
		db: db,
	}
}
