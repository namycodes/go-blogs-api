package migrations

import (
	"log"
	"com.namycodes/internal/blog"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB){
   err := db.AutoMigrate(
	blog.Blog{},
   )

   if err != nil {
	log.Fatalf("Migrations failed%s", err)
   }
}