package initializers

import "go_playground/go-jwt/models"

func SyncDatabase() {
	//DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.User{})
}
