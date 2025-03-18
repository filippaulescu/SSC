package main

import (
	"ic-project/controllers"
	"ic-project/initializers"

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

    r.Run() // Pornirea serverului pe portul implicit 8080
}
