package main

import (
	"github.com/bilguunerkh/unetei.mn/initializers"
	"github.com/bilguunerkh/unetei.mn/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Litter{})
}
