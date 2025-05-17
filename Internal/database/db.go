package database

import (
	"Pet-project/internal/taskService"
	"Pet-project/internal/userService"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (DB *gorm.DB) {
	dsn := "host=localhost user=testuser password=testpass dbname=testdb port=5433 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = DB.AutoMigrate(&userService.User{}, &taskService.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	return
}
