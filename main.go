package main

import (
	"github.com/bilguunerkh/unetei.mn/controllers"
	"github.com/bilguunerkh/unetei.mn/initializers"
	"github.com/bilguunerkh/unetei.mn/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()
	r.POST("/litters", controllers.LitterCreate)
	r.GET("/litters", controllers.GetLitters)
	r.GET("/litters/:id", controllers.GetLitter)
	r.PUT("/litters/:id", controllers.UpdateLitter)
	r.DELETE("/litters/:id", controllers.DeleteLitter)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
