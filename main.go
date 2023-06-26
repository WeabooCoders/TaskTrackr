package main

import (
	"fmt"

	"github.com/AvinFajarF/handlers"
	"github.com/AvinFajarF/initializers"
	"github.com/AvinFajarF/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.Migrate()
}

func main() {
	fmt.Println("Oke")

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/signup", handlers.SignUp)
	v1.POST("/sigin", handlers.SignIn)
	v1.GET("/coba",middleware.AuthMiddleware , handlers.Coba)

	router.Run(":8081")
}
