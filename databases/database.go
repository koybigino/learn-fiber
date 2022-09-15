package databases

import (
	"fmt"
	"github/koybigino/getting-started-fiber/models"

	// "log"
	"os"
	"strconv"

	// "github.com/joho/godotenv"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", os.Getenv("DB_HOSTNAME"), os.Getenv("DB_USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"), port)

	// dsn := "host=localhost user=postgres password=Bielem@*01 dbname=fiber_bd port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        dsn,
	}))

	if err != nil {
		fmt.Printf("Error to connect to our databse : %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("connection to the database ok .... %v", db)

	db.AutoMigrate(&models.User{}, &models.Post{})
	return db
}
