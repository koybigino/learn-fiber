package databases

import (
	"fmt"
	"github/koybigino/getting-started-fiber/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	dsn := "host=localhost user=postgres password=Bielem@*01 dbname=fiber_bd port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error to connect to our databse : %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("connection to the database ok .... %v", db)

	db.AutoMigrate(&models.Post{}, &models.User{})
	return db
}
