package repository

import (
	"fmt"
	"log"

	"github.com/Tes-sudo/online-learning-platform/user-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "user=yourusername dbname=userdb sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Successfully connected to database")

	// Auto Migrate the schema
	err = DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		log.Fatal("Failed to auto migrate:", err)
	}
}
