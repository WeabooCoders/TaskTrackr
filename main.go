package main

import (
	"github.com/AvinFajarF/handlers"
	"github.com/AvinFajarF/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.POST("/signup", handlers.SignUp)
	v1.POST("/sigin", handlers.SignIn)
	v1.GET("/coba", middleware.AuthMiddleware, handlers.Coba)
	v1.POST("/task", middleware.AuthMiddleware, handlers.CreateTask)
	v1.GET("/task", middleware.AuthMiddleware, handlers.GetTaskAllById)
	v1.GET("/task/:title", middleware.AuthMiddleware, handlers.FindTaskByTitle)
	v1.PUT("/task/update/:title", middleware.AuthMiddleware, handlers.UpdateTaskByTitle)

	router.Run(":8081")
}
