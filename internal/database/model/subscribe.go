package model

import (
	"time"

	"github.com/google/uuid"
)

// make model for subscribe user to package
type Subscribe struct {
	UserID        uuid.UUID `gorm:"column:user_id" json:"user_id"`
	PackageID     uuid.UUID `gorm:"column:package_id" json:"package_id"`
	PurchaseDate  time.Time `gorm:"column:purchase_date" json:"purchase_date"`
	ExpiredDate   time.Time `gorm:"column:expired_date" json:"expired_date"`
	PaymentMethod string    `gorm:"column:payment_method" json:"payment_method"`
	User          User      `gorm:"foreignKey:UserID"`
	Package       Package   `gorm:"foreignKey:PackageID"`
	BaseModel
}
