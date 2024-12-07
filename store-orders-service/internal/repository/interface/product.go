package interfaces

import (
	"context"
	"order-service/internal/domain/order"
)

type OrderRepository interface {
	Create(ctx context.Context, entity order.Entity) (id string, err error)
	List(ctx context.Context) (res []order.Entity, err error)
	Get(ctx context.Context, id string) (res order.Entity, err error)
	Delete(ctx context.Context, id string) (err error)
	Update(ctx context.Context, id string, entity order.Entity) (err error)
	Search(ctx context.Context, filter, value string) (res []order.Entity, err error)
}
