package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "postgresql://postgres:P4ssw0rd!@localhost:5433/Assignment5"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database!")
		return nil, err
	}
	log.Println("Database connected successfully.")
	return db, nil
}
