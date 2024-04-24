package config

import (
	"learn-go/internal/database/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	} else {
		log.Println("Connected to database")
	}
	db.AutoMigrate(&model.User{}, &model.Resource{}, &model.Package{}, &model.PackageResource{}, &model.Subscribe{})
	return db
}
