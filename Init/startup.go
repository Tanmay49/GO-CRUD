package Init

import (
	"log"
	"main/Models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}
 
	DB.AutoMigrate(&Models.Post{})
}
