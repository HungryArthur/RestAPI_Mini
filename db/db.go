package db

import (
	"RestAPI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dbb *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	dbb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключится к базе данных: %v", err)
	}

	dbb.AutoMigrate(&models.Message{})
}
