package service

import (
	"context"
	"order-service/internal/domain/order"
	interfaces "order-service/internal/repository/interface"
	services "order-service/internal/service/interface"
)

type OrderService struct {
	orderRepository interfaces.OrderRepository
}

func NewOrderService(orderRepository interfaces.OrderRepository) services.OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (ps *OrderService) CreateOrder(ctx context.Context, req order.Request) (id string, err error) {
	data := order.Entity{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Pricing:   req.Pricing,
		Status:    req.Status,
	}
	id, err = ps.orderRepository.Create(ctx, data)
	return
}

func (ps *OrderService) ListOrders(ctx context.Context) (res []order.Response, err error) {
	data, err := ps.orderRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	res = order.ParseFromEntities(data)
	return
}

func (ps *OrderService) GetOrder(ctx context.Context, id string) (res order.Response, err error) {
	data, err := ps.orderRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = order.ParseFromEntity(data)
	return
}

func (ps *OrderService) DeleteOrder(ctx context.Context, id string) (err error) {
	err = ps.orderRepository.Delete(ctx, id)
	return
}

func (ps *OrderService) UpdateOrder(ctx context.Context, id string, req order.Request) (err error) {
	data := order.Entity{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Pricing:   req.Pricing,
		Status:    req.Status,
	}
	err = ps.orderRepository.Update(ctx, id, data)
	return
}

func (ps *OrderService) SearchOrder(ctx context.Context, filter, value string) (res []order.Response, err error) {
	if !order.IsValidFilter(filter) || value == "" {
		err = order.ErrorInvalidSearch
		return
	}
	data, err := ps.orderRepository.Search(ctx, filter, value)
	if err != nil {
		return
	}
	res = order.ParseFromEntities(data)
	return
}
