package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"order-service/internal/domain/order"
	interfaces "order-service/internal/repository/interface"
	"strings"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) interfaces.OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (pr *OrderRepository) Create(ctx context.Context, data order.Entity) (id string, err error) {
	query := `
		INSERT INTO orders (user_id, product_id, pricing, status)
		VALUES ($1, $2, $3, $4) RETURNING id;`
	args := []any{
		data.UserID,
		pq.Array(data.ProductID),
		data.Pricing,
		data.Status,
	}
	if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = order.ErrorNotFound
		}
	}
	return
}

func (pr *OrderRepository) List(ctx context.Context) (projects []order.Entity, err error) {
	query := `SELECT * FROM orders ORDER BY id;`
	err = pr.db.SelectContext(ctx, &projects, query)
	return
}

func (pr *OrderRepository) Get(ctx context.Context, id string) (dest order.Entity, err error) {
	query := `SELECT * FROM orders WHERE id = $1;`
	args := []any{id}
	err = pr.db.GetContext(ctx, &dest, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = order.ErrorNotFound
		}
	}
	return
}

func (pr *OrderRepository) Delete(ctx context.Context, id string) (err error) {
	query := `DELETE FROM orders WHERE id = $1 RETURNING id;`
	args := []any{id}
	if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = order.ErrorNotFound
		}
	}
	return
}

func (pr *OrderRepository) Update(ctx context.Context, id string, data order.Entity) (err error) {
	sets, args := pr.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, id)
		query := fmt.Sprintf("UPDATE orders SET %s WHERE id = $%d RETURNING id;", strings.Join(sets, ","), len(args))
		if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = order.ErrorNotFound
			}
		}
	}
	return
}

func (pr *OrderRepository) Search(ctx context.Context, filter, value string) (dest []order.Entity, err error) {
	dest = []order.Entity{}
	filter = pr.prepareFilter(filter)
	query := fmt.Sprintf("SELECT * FROM orders WHERE %s = $1", filter)
	err = pr.db.SelectContext(ctx, &dest, query, value)
	if err != nil {
		return
	}
	if len(dest) == 0 {
		err = order.ErrorNotFound
		return
	}
	return
}

func (pr *OrderRepository) prepareArgs(data order.Entity) (sets []string, args []any) {
	if data.UserID != "" {
		sets = append(sets, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, data.UserID)
	}
	if len(data.ProductID) != 0 {
		sets = append(sets, fmt.Sprintf("product_id = $%d", len(args)+1))
		args = append(args, pq.Array(data.ProductID))
	}
	if data.Pricing != 0 {
		sets = append(sets, fmt.Sprintf("pricing = $%d", len(args)+1))
		args = append(args, data.Pricing)
	}
	if data.Status != "" {
		sets = append(sets, fmt.Sprintf("status = $%d", len(args)+1))
		args = append(args, data.Status)

	}
	return
}

func (pr *OrderRepository) prepareFilter(arg string) string {
	switch arg {
	case "user_id":
		return "user_id"
	case "status":
		return "status"
	default:
		return ""
	}
}
