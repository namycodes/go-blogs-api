package migrations

import (
	"log"
	"com.namycodes/internal/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB){
   err := db.AutoMigrate(
	models.Blog{},
	models.User{},
   )

   if err != nil {
	log.Fatalf("Migrations failed%s", err)
   }
}