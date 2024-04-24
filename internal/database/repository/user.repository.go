package repository

import (
	"learn-go/internal/database/model"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(args ...int) ([]model.User, error)
	FindOne(id uuid.UUID) (model.User, error)
	FindOneByEmail(email string) (model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(user *model.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r *UserRepositoryImpl) FindAll(args ...int) ([]model.User, error) {
	var users []model.User
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
	err := query.Find(&users).Error
	// count,
	if err != nil {
		return nil, err
	}
	return users, nil
}

// find by id
func (r *UserRepositoryImpl) FindOne(id uuid.UUID) (model.User, error) {
	var user model.User
	err := r.db.Where("id =?", id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindOneByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Unscoped().Where("email =?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// create new user
func (r *UserRepositoryImpl) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// update user
func (r *UserRepositoryImpl) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// delete user
func (r *UserRepositoryImpl) Delete(user *model.User) error {
	return r.db.Delete(user).Error
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}
