package database

import (
	"fmt"
	"log"

	"ticketing-application/internal/config"
	"ticketing-application/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "postgres")
	password := config.GetEnv("DB_PASSWORD", "secret")
	dbname := config.GetEnv("DB_NAME", "postgres")
	sslmode := config.GetEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	err = DB.AutoMigrate(&model.Slot{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	seedInitialSlots()
}

func seedInitialSlots() {
	var count int64
	DB.Model(&model.Slot{}).Count(&count)
	if count == 0 {
		slots := []model.Slot{
			{TimeSlot: "08:00 AM - 09:00 AM", Status: "available"},
			{TimeSlot: "09:00 AM - 10:00 AM", Status: "available"},
			{TimeSlot: "10:00 AM - 11:00 AM", Status: "available"},
			{TimeSlot: "11:00 AM - 12:00 PM", Status: "available"},
		}
		DB.Create(&slots)
		fmt.Println("Database seeded with initial available slots!")
	}
}
