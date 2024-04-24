package model

// make role for user
type Role struct {
	Name string `gorm:"column:name" json:"name"`
	BaseModel
}
