package model

// make package for subscription
type Package struct {
	Name        string  `gorm:"column:name" json:"name"`
	Image       string  `gorm:"column:image" json:"image"`
	Price       float64 `gorm:"column:price" json:"price"`
	Description string  `gorm:"column:description" json:"description"`
	BaseModel

	// package have many subscribe
	Subscribes []Subscribe `gorm:"foreignKey:PackageID"`
	// package has many resource
	PackageResource []PackageResource
}
