package routes

import (
	"github.com/gin-gonic/gin"
	"payment-service/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, paymentHandler *handler.PaymentHandler) {

	router.GET("/", paymentHandler.ListPayments)
	router.POST("/", paymentHandler.CreatePayment)
	router.GET("/:id", paymentHandler.GetPayment)
	router.PUT("/:id", paymentHandler.UpdatePayment)
	router.DELETE("/:id", paymentHandler.DeletePayment)
	router.GET("/search", paymentHandler.SearchPayments)

}
