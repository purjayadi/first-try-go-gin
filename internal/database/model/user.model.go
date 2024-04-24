package model

type User struct {
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email;unique" json:"email"`
	Password string `gorm:"column:password;<-:false" json:"password" `
	BaseModel

	// user have many subscribe
	Subscribes []Subscribe `gorm:"foreignKey:UserID"`
}
