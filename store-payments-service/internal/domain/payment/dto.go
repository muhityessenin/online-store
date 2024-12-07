package payment

import (
	"errors"
	"time"
)

var (
	ErrorNotFound            = errors.New("payment not found")
	ErrorInvalidDate         = errors.New("invalid date format")
	ErrorFailedToMakePayment = errors.New("failed to make payment")
	ErrorInvalidAmount       = errors.New("invalid amount")
	ErrorInvalidUserID       = errors.New("invalid user id")
	ErrorInvalidOrderID      = errors.New("invalid order id")
)

type Request struct {
	UserID  string  `json:"user_id"`
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
}

type Response struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	OrderID   string    `json:"order_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func ParseFromEntity(entity Entity) Response {
	return Response{
		ID:        entity.ID,
		UserID:    entity.UserID,
		OrderID:   entity.OrderID,
		Amount:    entity.Amount,
		Status:    entity.Status,
		CreatedAt: entity.CreatedAt,
	}
}

func ParseFromEntities(data []Entity) (res []Response) {
	res = make([]Response, 0)
	for _, entity := range data {
		res = append(res, ParseFromEntity(entity))
	}
	return
}

func (r *Request) Validate() error {
	if r.UserID == "" {
		return ErrorInvalidUserID
	}
	if r.OrderID == "" {
		return ErrorInvalidOrderID
	}
	if r.Amount <= 0 {
		return ErrorInvalidAmount
	}
	return nil
}
