package handler

import (
	"api-gateway-service/pkg/response"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type ProductHandler struct {
	productUrl string
}

func NewProductHandler(productUrl string) *ProductHandler {
	return &ProductHandler{productUrl}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body product.Request true "Product data"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products [post]
func (p *ProductHandler) CreateProduct(c *gin.Context) {
	req, err := http.NewRequest("POST", p.productUrl, c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(c.Writer, resp.Body)
}

// ListProducts godoc
// @Summary List all products
// @Description List all products
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products [get]
func (p *ProductHandler) ListProducts(c *gin.Context) {
	req, err := http.NewRequest("GET", p.productUrl, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)
}

// GetProduct godoc
// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products/{id} [get]
func (p *ProductHandler) GetProduct(c *gin.Context) {
	req, err := http.NewRequest("GET", p.productUrl+"/"+c.Param("id"), nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)
}

// UpdateProduct godoc
// @Summary Update product by ID
// @Description Update product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Param product body product.Request true "Product data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products/{id} [put]
func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	req, err := http.NewRequest("PUT", p.productUrl+"/"+c.Param("id"), c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)
}

// DeleteProduct godoc
// @Summary Delete product by ID
// @Description Delete product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products/{id} [delete]
func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	req, err := http.NewRequest("DELETE", p.productUrl+"/"+c.Param("id"), nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)
}

// SearchProducts godoc
// @Summary Search products
// @Description Search products
// @Tags products
// @Accept  json
// @Produce  json
// @Param filter query string false "Search filter"
// @Param value query string false "Search value"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products/search [get]
func (p *ProductHandler) SearchProducts(c *gin.Context) {
	req, err := http.NewRequest("GET", p.productUrl+"/search?filter="+c.Query("filter")+"&value="+c.Query("value"), nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)
}
