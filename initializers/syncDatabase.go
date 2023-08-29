package initializers

import "github.com/bilguunerkh/unetei.mn/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
