package interfaces

import (
	"context"
	"order-service/internal/domain/order"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req order.Request) (id string, err error)
	ListOrders(ctx context.Context) (res []order.Response, err error)
	GetOrder(ctx context.Context, id string) (res order.Response, err error)
	DeleteOrder(ctx context.Context, id string) (err error)
	UpdateOrder(ctx context.Context, id string, req order.Request) (err error)
	SearchOrder(ctx context.Context, filter, value string) (res []order.Response, err error)
}
