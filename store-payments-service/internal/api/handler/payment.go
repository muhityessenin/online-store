package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-service/internal/domain/payment"
	interfaces "payment-service/internal/service/interface"
	"payment-service/pkg/response"
)

type PaymentHandler struct {
	paymentService interfaces.PaymentService
}

func NewPaymentHandler(service interfaces.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: service,
	}
}

// CreatePayment godoc
// @Summary Create a new payment
// @Description Create a new payment with the input payload
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body payment.Request true "Payment Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /payments [post]
func (th *PaymentHandler) CreatePayment(c *gin.Context) {
	req := payment.Request{}
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

	res, err := th.paymentService.CreatePayment(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, payment.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields must be unique", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		if errors.Is(err, payment.ErrorInvalidDate) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to create payment", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "the payment was successfully created", res, nil)
	c.JSON(http.StatusCreated, successRes)
}

// ListPayments godoc
// @Summary List all payments
// @Description Get a list of payments
// @Tags payments
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /payments [get]
func (th *PaymentHandler) ListPayments(c *gin.Context) {
	res, err := th.paymentService.ListPayments(c.Request.Context())
	if err != nil {
		if errors.Is(err, payment.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no tasks found", "", nil)
			c.JSON(http.StatusOK, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to list tasks", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the tasks list", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetPayment godoc
// @Summary Get a payment by ID
// @Description Get details of a payment by its ID
// @Tags payments
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /payments/{id} [get]
func (th *PaymentHandler) GetPayment(c *gin.Context) {
	id := c.Param("id")
	res, err := th.paymentService.GetPayment(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, payment.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no tasks found", "", nil)
			c.JSON(http.StatusOK, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to get payment", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the payment details", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// UpdatePayment godoc
// @Summary Update a payment by ID
// @Description Update details of a payment by its ID
// @Tags payments
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Param payment body payment.Request true "Task Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /payments/{id} [put]
func (th *PaymentHandler) UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	req := payment.Request{}
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

	err := th.paymentService.UpdatePayment(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, payment.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "payment not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return

		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to update payment", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the payment was successfully updated", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// DeletePayment godoc
// @Summary Delete a payment by ID
// @Description Delete a payment by its ID
// @Tags payments
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /payments/{id} [delete]
func (th *PaymentHandler) DeletePayment(c *gin.Context) {
	id := c.Param("id")
	err := th.paymentService.DeletePayment(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, payment.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "payment not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to delete payment", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the payment was successfully deleted", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// SearchPayments godoc
// @Summary Search payments
// @Description Search payments by filter and value
// @Tags payments
// @Produce json
// @Param filter query string true "Filter"
// @Param value query string true "Value"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /payments/search [get]
func (th *PaymentHandler) SearchPayments(c *gin.Context) {
	filter := c.Query("filter")
	value := c.Query("value")
	res, err := th.paymentService.SearchPayments(c.Request.Context(), filter, value)
	if err != nil {
		if errors.Is(err, payment.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no tasks found", "", nil)
			c.JSON(http.StatusOK, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to search payments", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the tasks list", res, nil)
	c.JSON(http.StatusOK, successRes)
}
