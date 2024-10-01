package initializers

import "github.com/taco-tortilla/jwt-go/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
