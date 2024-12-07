package routes

import (
	"github.com/gin-gonic/gin"
	"order-service/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, orderHandler *handler.OrderHandler) {
	router.GET("/", orderHandler.ListOrders)
	router.POST("/", orderHandler.CreateOrder)
	router.GET("/:id", orderHandler.GetOrder)
	router.PUT("/:id", orderHandler.UpdateOrder)
	router.DELETE("/:id", orderHandler.DeleteOrder)
	router.GET("/search", orderHandler.SearchOrders)

}
