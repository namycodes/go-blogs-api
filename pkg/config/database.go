package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB{
	datasourse := "host=localhost user=postgres password=B1u3_r0s3@ dbname=gousers port=5432 sslmode=disable"

	var err error

	db, err := gorm.Open(postgres.Open(datasourse), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to DB %s", err)
	}

	fmt.Println("Connected To Database Successfully")

	return db
}