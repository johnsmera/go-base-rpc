package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/johnsmera/chall/domain"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("gorm_dsn")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	// defer db.Close()

	db.AutoMigrate(&domain.User{})

	return db
}
