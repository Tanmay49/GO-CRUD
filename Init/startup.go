package Init

import (
	"log"
	"main/Models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&Models.Post{})
}
