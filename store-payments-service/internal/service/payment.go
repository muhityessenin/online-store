package service

import (
	"context"
	"fmt"
	"log"
	"payment-service/internal/domain/payment"
	interfaces "payment-service/internal/repository/interface"
	services "payment-service/internal/service/interface"
)

type PaymentService struct {
	paymentRepository interfaces.PaymentRepository
}

func NewPaymentService(repository interfaces.PaymentRepository) services.PaymentService {
	return &PaymentService{
		paymentRepository: repository,
	}
}

func (ts *PaymentService) CreatePayment(ctx context.Context, req payment.Request) (id string, err error) {
	_, err = MakePayment()
	var status string
	if err != nil {
		status = "failed"
		log.Printf("failed to make payment: %v", fmt.Sprintf("%v", err))
		err = payment.ErrorFailedToMakePayment
	} else {
		status = "success"
	}
	data := payment.Entity{
		UserID:  req.UserID,
		OrderID: req.OrderID,
		Amount:  req.Amount,
		Status:  status,
	}
	id, err = ts.paymentRepository.Create(ctx, data)
	return
}

func (ts *PaymentService) ListPayments(ctx context.Context) (res []payment.Response, err error) {
	data, err := ts.paymentRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	res = payment.ParseFromEntities(data)
	return
}

func (ts *PaymentService) GetPayment(ctx context.Context, id string) (res payment.Response, err error) {
	data, err := ts.paymentRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = payment.ParseFromEntity(data)
	return
}

func (ts *PaymentService) DeletePayment(ctx context.Context, id string) (err error) {
	err = ts.paymentRepository.Delete(ctx, id)
	return
}

func (ts *PaymentService) UpdatePayment(ctx context.Context, id string, req payment.Request) (err error) {
	data := payment.Entity{
		UserID:  req.UserID,
		OrderID: req.OrderID,
		Amount:  req.Amount,
	}
	err = ts.paymentRepository.Update(ctx, id, data)
	return
}

func (ts *PaymentService) SearchPayments(ctx context.Context, filter, value string) (res []payment.Response, err error) {
	data, err := ts.paymentRepository.Search(ctx, filter, value)
	if err != nil {
		return nil, err
	}
	res = payment.ParseFromEntities(data)
	return
}
