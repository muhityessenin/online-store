package routes

import (
	"github.com/gin-gonic/gin"
	"product-service/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, productHandler *handler.ProductHandler) {
	router.GET("/", productHandler.ListProducts)
	router.POST("/", productHandler.CreateProduct)
	router.GET("/:id", productHandler.GetProduct)
	router.PUT("/:id", productHandler.UpdateProduct)
	router.DELETE("/:id", productHandler.DeleteProduct)
	router.GET("/search", productHandler.SearchProduct)

}
