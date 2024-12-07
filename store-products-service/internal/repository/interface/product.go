package interfaces

import (
	"context"
	"product-service/internal/domain/product"
)

type ProductRepository interface {
	Create(ctx context.Context, entity product.Entity) (id string, err error)
	List(ctx context.Context) (res []product.Entity, err error)
	Get(ctx context.Context, id string) (res product.Entity, err error)
	Delete(ctx context.Context, id string) (err error)
	Update(ctx context.Context, id string, entity product.Entity) (err error)
	Search(ctx context.Context, filter, value string) (res []product.Entity, err error)
}
