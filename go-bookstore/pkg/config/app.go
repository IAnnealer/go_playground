package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dsn := "host=arjuna.db.elephantsql.com user=wtpmedwj password=f3_hXc1ZO1Wla0jyLuELJY7iDNvxhz3A dbname=wtpmedwj port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect database")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
