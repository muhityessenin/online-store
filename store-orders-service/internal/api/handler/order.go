package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/internal/domain/order"
	interfaces "order-service/internal/service/interface"
	"order-service/pkg/response"
)

type OrderHandler struct {
	orderService interfaces.OrderService
}

func NewOrderHandler(service interfaces.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: service,
	}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input payload
// @Tags orders
// @Accept json
// @Produce json
// @Param payment body order.Request true "Order Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders [post]
func (th *OrderHandler) CreateOrder(c *gin.Context) {
	req := order.Request{}
	if err := c.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := th.orderService.CreateOrder(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, order.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields must be unique", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to create order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "the order was successfully created", res, nil)
	c.JSON(http.StatusCreated, successRes)
}

// ListOrders godoc
// @Summary List all orders
// @Description Get a list of orders
// @Tags orders
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders [get]
func (th *OrderHandler) ListOrders(c *gin.Context) {
	res, err := th.orderService.ListOrders(c.Request.Context())
	if err != nil {
		if errors.Is(err, order.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no orders found", "", nil)
			c.JSON(http.StatusOK, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to list orders", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the orders list", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetOrder godoc
// @Summary Get a order by ID
// @Description Get details of a order by its ID
// @Tags orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders/{id} [get]
func (th *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	res, err := th.orderService.GetOrder(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, order.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no orders found", "", nil)
			c.JSON(http.StatusOK, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to get order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the order details", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// UpdateOrder godoc
// @Summary Update a order by ID
// @Description Update details of a order by its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body order.Request true "Order Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders/{id} [put]
func (th *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	req := order.Request{}
	if err := c.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := th.orderService.UpdateOrder(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, order.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "order not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return

		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to update order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the order was successfully updated", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// DeleteOrder godoc
// @Summary Delete a order by ID
// @Description Delete a order by its ID
// @Tags orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders/{id} [delete]
func (th *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := th.orderService.DeleteOrder(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, order.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "order not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to delete order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the order was successfully deleted", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// SearchOrders godoc
// @Summary Search orders by filter
// @Description Search orders by filter
// @Tags orders
// @Produce json
// @Param filter query string true "Filter"
// @Param value query string true "Value"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /orders/search [get]
func (th *OrderHandler) SearchOrders(c *gin.Context) {
	filter := c.Query("filter")
	value := c.Query("value")
	res, err := th.orderService.SearchOrder(c.Request.Context(), filter, value)
	if err != nil {
		if errors.Is(err, order.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no orders found", nil, nil)
			c.JSON(http.StatusOK, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to search orders", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the orders list", res, nil)
	c.JSON(http.StatusOK, successRes)
}
