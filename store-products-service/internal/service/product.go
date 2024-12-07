package service

import (
	"context"
	"product-service/internal/domain/product"
	interfaces "product-service/internal/repository/interface"
	services "product-service/internal/service/interface"
)

type ProductService struct {
	productRepository interfaces.ProductRepository
}

func NewProductService(productRepository interfaces.ProductRepository) services.ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (ps *ProductService) CreateProduct(ctx context.Context, req product.Request) (id string, err error) {
	data := product.Entity{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Quantity:    req.Quantity,
	}
	id, err = ps.productRepository.Create(ctx, data)
	return
}

func (ps *ProductService) ListProduct(ctx context.Context) (res []product.Response, err error) {
	data, err := ps.productRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	res = product.ParseFromEntities(data)
	return
}

func (ps *ProductService) GetProduct(ctx context.Context, id string) (res product.Response, err error) {
	data, err := ps.productRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = product.ParseFromEntity(data)
	return
}

func (ps *ProductService) DeleteProduct(ctx context.Context, id string) (err error) {
	err = ps.productRepository.Delete(ctx, id)
	return
}

func (ps *ProductService) UpdateProduct(ctx context.Context, id string, req product.Request) (err error) {
	data := product.Entity{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Quantity:    req.Quantity,
	}
	err = ps.productRepository.Update(ctx, id, data)
	return
}

func (ps *ProductService) SearchProduct(ctx context.Context, filter, value string) (res []product.Response, err error) {
	if !product.IsValidFilter(filter) || value == "" {
		err = product.ErrorInvalidSearch
		return
	}
	data, err := ps.productRepository.Search(ctx, filter, value)
	if err != nil {
		return
	}
	res = product.ParseFromEntities(data)
	return
}
