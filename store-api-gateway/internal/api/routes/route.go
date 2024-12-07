package routes

import (
	"api-gateway-service/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, orderHandler *handler.OrderHandler, productHandler *handler.ProductHandler, paymentHandler *handler.PaymentHandler) {
	users := router.Group("/users")
	{
		users.GET("/", userHandler.ListUsers)
		users.POST("/", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
		users.PUT("/search", userHandler.SearchUser)
	}

	products := router.Group("/products")
	{
		products.GET("/", productHandler.ListProducts)
		products.POST("/", productHandler.CreateProduct)
		products.GET("/:id", productHandler.GetProduct)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
		products.PUT("/search", productHandler.SearchProducts)
	}

	orders := router.Group("/orders")
	{
		orders.GET("/", orderHandler.ListOrders)
		orders.POST("/", orderHandler.CreateOrder)
		orders.GET("/:id", orderHandler.GetOrder)
		orders.PUT("/:id", orderHandler.UpdateOrder)
		orders.DELETE("/:id", orderHandler.DeleteOrder)
		orders.PUT("/search", orderHandler.SearchOrders)
	}

	payments := router.Group("/payments")
	{
		payments.GET("/", paymentHandler.ListPayments)
		payments.POST("/", paymentHandler.CreatePayment)
		payments.GET("/:id", paymentHandler.GetPayment)
		payments.PUT("/:id", paymentHandler.UpdatePayment)
		payments.DELETE("/:id", paymentHandler.DeletePayment)
		payments.PUT("/search", paymentHandler.SearchPayments)
	}
}
