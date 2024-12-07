package interfaces

import (
	"context"
	"product-service/internal/domain/product"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req product.Request) (id string, err error)
	ListProduct(ctx context.Context) (res []product.Response, err error)
	GetProduct(ctx context.Context, id string) (res product.Response, err error)
	DeleteProduct(ctx context.Context, id string) (err error)
	UpdateProduct(ctx context.Context, id string, req product.Request) (err error)
	SearchProduct(ctx context.Context, filter, value string) (res []product.Response, err error)
}
