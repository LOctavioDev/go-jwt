package main

import (
	"jwt-go/controllers"
	"jwt-go/initializers"
	"jwt-go/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/singup", controllers.Singup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiereAuth, controllers.Validate)
	r.POST("/logout", controllers.Logout)

	r.Run()
}
