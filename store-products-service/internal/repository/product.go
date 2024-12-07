package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"product-service/internal/domain/product"
	interfaces "product-service/internal/repository/interface"
	"strings"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) interfaces.ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) Create(ctx context.Context, data product.Entity) (id string, err error) {
	query := `
		INSERT INTO products (title, description, price, category, quantity)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	args := []any{
		data.Title,
		data.Description,
		data.Price,
		data.Category,
		data.Quantity,
	}
	if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = product.ErrorNotFound
		}
	}
	return
}

func (pr *ProductRepository) List(ctx context.Context) (projects []product.Entity, err error) {
	query := `SELECT * FROM products ORDER BY id;`
	err = pr.db.SelectContext(ctx, &projects, query)
	return
}

func (pr *ProductRepository) Get(ctx context.Context, id string) (dest product.Entity, err error) {
	query := `SELECT * FROM products WHERE id = $1;`
	args := []any{id}
	err = pr.db.GetContext(ctx, &dest, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = product.ErrorNotFound
		}
	}
	return
}

func (pr *ProductRepository) Delete(ctx context.Context, id string) (err error) {
	query := `DELETE FROM products WHERE id = $1 RETURNING id;`
	args := []any{id}
	if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = product.ErrorNotFound
		}
	}
	return
}

func (pr *ProductRepository) Update(ctx context.Context, id string, data product.Entity) (err error) {
	sets, args := pr.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, id)
		query := fmt.Sprintf("UPDATE products SET %s WHERE id = $%d RETURNING id;", strings.Join(sets, ","), len(args))
		if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = product.ErrorNotFound
			}
		}
	}
	return
}

func (pr *ProductRepository) Search(ctx context.Context, filter, value string) (dest []product.Entity, err error) {
	dest = []product.Entity{}
	filter = pr.prepareFilter(filter)
	query := fmt.Sprintf("SELECT * FROM products WHERE %s = $1", filter)
	err = pr.db.SelectContext(ctx, &dest, query, value)
	if err != nil {
		return
	}
	if len(dest) == 0 {
		err = product.ErrorNotFound
		return
	}
	return
}

func (pr *ProductRepository) prepareArgs(data product.Entity) (sets []string, args []any) {
	if data.Title != "" {
		args = append(args, data.Title)
		sets = append(sets, fmt.Sprintf("title = $%d", len(args)))
	}
	if data.Description != "" {
		args = append(args, data.Description)
		sets = append(sets, fmt.Sprintf("description = $%d", len(args)))
	}
	if data.Price != 0 {
		args = append(args, data.Price)
		sets = append(sets, fmt.Sprintf("price = $%d", len(args)))
	}
	if data.Category != "" {
		args = append(args, data.Category)
		sets = append(sets, fmt.Sprintf("category = $%d", len(args)))
	}
	if data.Quantity != 0 {
		args = append(args, data.Quantity)
		sets = append(sets, fmt.Sprintf("quantity = $%d", len(args)))
	}
	return
}

func (pr *ProductRepository) prepareFilter(filter string) string {
	switch filter {
	case "title":
		return "title"
	case "category":
		return "category"
	default:
		return ""
	}
}
