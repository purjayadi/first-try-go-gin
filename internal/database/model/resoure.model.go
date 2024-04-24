package model

// make resource for package subscription
type Resource struct {
	Name            string `gorm:"column:name" json:"name"`
	PackageResource []PackageResource
	BaseModel
}
