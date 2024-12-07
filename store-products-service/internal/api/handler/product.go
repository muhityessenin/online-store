package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"product-service/internal/domain/product"
	interfaces "product-service/internal/service/interface"
	"product-service/pkg/response"
)

type ProductHandler struct {
	productService interfaces.ProductService
}

func NewProductHandler(service interfaces.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}
}

// CreateProduct godoc
// @Summary Create a new order
// @Description Create a new order with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param payment body product.Request true "Product Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products [post]
func (th *ProductHandler) CreateProduct(c *gin.Context) {
	req := product.Request{}
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

	res, err := th.productService.CreateProduct(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, product.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields must be unique", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		if errors.Is(err, product.ErrorInvalidDate) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
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

// ListProducts godoc
// @Summary List all products
// @Description Get a list of products
// @Tags products
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products [get]
func (th *ProductHandler) ListProducts(c *gin.Context) {
	res, err := th.productService.ListProduct(c.Request.Context())
	if err != nil {
		if errors.Is(err, product.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no products found", "", nil)
			c.JSON(http.StatusOK, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to list products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the products list", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetProduct godoc
// @Summary Get a order by ID
// @Description Get details of a order by its ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products/{id} [get]
func (th *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	res, err := th.productService.GetProduct(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, product.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no products found", "", nil)
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

// UpdateProduct godoc
// @Summary Update a order by ID
// @Description Update details of a order by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param order body product.Request true "Product Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products/{id} [put]
func (th *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	req := product.Request{}
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

	err := th.productService.UpdateProduct(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, product.ErrorNotFound) {
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

// DeleteProduct godoc
// @Summary Delete a order by ID
// @Description Delete a order by its ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products/{id} [delete]
func (th *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := th.productService.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, product.ErrorNotFound) {
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

// SearchProduct godoc
// @Summary Search for a order by filter
// @Description Search for a order by filter
// @Tags products
// @Produce json
// @Param filter query string true "Filter"
// @Param value query string true "Value"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products/search [get]
func (th *ProductHandler) SearchProduct(c *gin.Context) {
	filter := c.Query("filter")
	value := c.Query("value")
	res, err := th.productService.SearchProduct(c.Request.Context(), filter, value)
	if err != nil {
		if errors.Is(err, product.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no products found", "", nil)
			c.JSON(http.StatusOK, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to search products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the products list", res, nil)
	c.JSON(http.StatusOK, successRes)
}
