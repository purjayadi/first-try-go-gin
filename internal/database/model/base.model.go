package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;DeletedAt" json:"deleted_at"`
}

// BeforeCreate hook will be called before creating a new record
func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID for the ID field
	base.ID = uuid.New()
	return nil
}
