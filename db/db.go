package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // GORM needs this implicitly
	"github.com/jinzhu/gorm"
)

// New initializes GORM connection to database
func New() *gorm.DB {
	connectionString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Database connection failed: %s", err.Error())
	}

	return db
}
