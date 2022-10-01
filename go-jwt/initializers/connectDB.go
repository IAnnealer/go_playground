package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = DB

	if err != nil {
		panic("Failed to connect DB")
	}
}
