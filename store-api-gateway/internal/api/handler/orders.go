package handler

import (
	"api-gateway-service/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	orderUrl string
}

func NewOrderHandler(orderUrl string) *OrderHandler {
	return &OrderHandler{orderUrl}
}

// CreateOrder godoc
// @Summary Create order
// @Description Create order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body order.Request true "Order data"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders [post]
func (o *OrderHandler) CreateOrder(c *gin.Context) {
	req, err := http.NewRequest("POST", o.orderUrl, c.Request.Body)
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

// ListOrders godoc
// @Summary List all orders
// @Description List all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders [get]
func (o *OrderHandler) ListOrders(c *gin.Context) {
	req, err := http.NewRequest("GET", o.orderUrl, nil)
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

// GetOrder godoc
// @Summary Get order by ID
// @Description Get order by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders/{id} [get]
func (o *OrderHandler) GetOrder(c *gin.Context) {
	req, err := http.NewRequest("GET", o.orderUrl+"/"+c.Param("id"), nil)
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

// UpdateOrder godoc
// @Summary Update order by ID
// @Description Update order by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Param order body order.Request true "Order data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders/{id} [put]
func (o *OrderHandler) UpdateOrder(c *gin.Context) {
	req, err := http.NewRequest("PUT", o.orderUrl+"/"+c.Param("id"), c.Request.Body)
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

// DeleteOrder godoc
// @Summary Delete order by ID
// @Description Delete order by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Success 204 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders/{id} [delete]
func (o *OrderHandler) DeleteOrder(c *gin.Context) {
	req, err := http.NewRequest("DELETE", o.orderUrl+"/"+c.Param("id"), nil)
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

// SearchOrders godoc
// @Summary Search orders
// @Description Search orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Param filter query string true "Filter"
// @Param value query string true "Value"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders/search [get]
func (o *OrderHandler) SearchOrders(c *gin.Context) {
	req, err := http.NewRequest("GET", o.orderUrl+"/search?filter="+c.Query("filter")+"&value="+c.Query("value"), nil)
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
