package model

import "github.com/google/uuid"

// make relation resource with package
type PackageResource struct {
	PackageID  uuid.UUID `gorm:"column:package_id" json:"package_id"`
	ResourceID uuid.UUID `gorm:"column:resource_id" json:"resource_id"`
	Package    Package   `gorm:"foreignKey:PackageID"`
	Resource   Resource  `gorm:"foreignKey:ResourceID"`
	BaseModel
}
