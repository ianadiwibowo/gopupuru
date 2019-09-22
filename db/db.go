package db

import (
	"fmt"
	"log"
	"os"

	// go-sql-driver/mysql is needed implicitly by GORM
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SetupDatabase initializes GORM
func SetupDatabase() *gorm.DB {
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
