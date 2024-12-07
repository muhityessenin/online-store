package interfaces

import (
	"context"
	"users-service/internal/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, data user.Entity) (id string, err error)
	List(ctx context.Context) (users []user.Entity, err error)
	Get(ctx context.Context, id string) (dest user.Entity, err error)
	Delete(ctx context.Context, id string) (err error)
	Update(ctx context.Context, id string, data user.Entity) (err error)
	Search(ctx context.Context, filter, value string) (users []user.Entity, err error)
}
