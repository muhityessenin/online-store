package interfaces

import (
	"context"
	"payment-service/internal/domain/payment"
)

type PaymentRepository interface {
	Create(ctx context.Context, entity payment.Entity) (id string, err error)
	List(ctx context.Context) (res []payment.Entity, err error)
	Get(ctx context.Context, id string) (res payment.Entity, err error)
	Delete(ctx context.Context, id string) (err error)
	Update(ctx context.Context, id string, entity payment.Entity) (err error)
	Search(ctx context.Context, filter, value string) (res []payment.Entity, err error)
}
