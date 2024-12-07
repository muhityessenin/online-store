package interfaces

import (
	"context"
	"payment-service/internal/domain/payment"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, req payment.Request) (id string, err error)
	ListPayments(ctx context.Context) (res []payment.Response, err error)
	GetPayment(ctx context.Context, id string) (res payment.Response, err error)
	DeletePayment(ctx context.Context, id string) (err error)
	UpdatePayment(ctx context.Context, id string, req payment.Request) (err error)
	SearchPayments(ctx context.Context, filter, value string) (res []payment.Response, err error)
}
