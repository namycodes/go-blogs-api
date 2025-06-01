package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB{
	datasourse := fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"),
)
	var err error

	db, err := gorm.Open(postgres.Open(datasourse), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to DB %s", err)
	}

	fmt.Println("Connected To Database Successfully")

	return db
}