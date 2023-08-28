package main

import (
	"github.com/bilguunerkh/unetei.mn/controllers"
	"github.com/bilguunerkh/unetei.mn/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/litters", controllers.LitterCreate)
	r.GET("/litters", controllers.GetLitters)
	r.GET("/litters/:id", controllers.GetLitter)
	r.PUT("/litters/:id", controllers.UpdateLitter)
	r.DELETE("/litters/:id", controllers.DeleteLitter)
	r.Run()
}
