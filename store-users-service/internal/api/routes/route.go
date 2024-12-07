package routes

import (
	"github.com/gin-gonic/gin"
	"users-service/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {
	router.GET("/", userHandler.ListUsers)
	router.POST("/", userHandler.CreateUser)
	router.GET("/:id", userHandler.GetUser)
	router.PUT("/:id", userHandler.UpdateUser)
	router.DELETE("/:id", userHandler.DeleteUser)
	router.GET("/search", userHandler.SearchUsers)
}
