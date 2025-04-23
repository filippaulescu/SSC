package main

import (
	"ic-project/controllers"
	"ic-project/initializers"
	"ic-project/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDb()
    initializers.SyncDatabase()
}

func main() {
    r := gin.Default()

    // Corectat
    r.POST("/signup", controllers.Signup)
    r.POST("/login", controllers.Login)
    r.GET("/validate", middleware.RequireAuth, controllers.Validate)
    r.Run(":3000") // Rulează explicit pe portul 3000
}
