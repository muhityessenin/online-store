package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"payment-service/internal/domain/payment"
	"strconv"
	"strings"
)

type PaymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (pr *PaymentRepository) Create(ctx context.Context, data payment.Entity) (id string, err error) {
	query := `INSERT INTO payments (user_id, order_id, amount, status) VALUES ($1, $2, $3, $4) RETURNING id;`
	args := []any{
		data.UserID,
		data.OrderID,
		data.Amount,
		data.Status,
	}
	if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = payment.ErrorNotFound
		}
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			return "", payment.ErrorNotFound
		}
	}
	return
}

func (pr *PaymentRepository) Delete(ctx context.Context, id string) (err error) {
	query := `DELETE FROM payments WHERE id = $1 RETURNING id;`
	args := []any{id}
	if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = payment.ErrorNotFound
			return
		}
	}

	return
}

func (pr *PaymentRepository) Get(ctx context.Context, id string) (dest payment.Entity, err error) {
	query := `SELECT * FROM payments WHERE id = $1;`
	args := []any{id}
	err = pr.db.GetContext(ctx, &dest, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		err = payment.ErrorNotFound
	}
	return
}

func (pr *PaymentRepository) List(ctx context.Context) (dest []payment.Entity, err error) {
	query := `SELECT * FROM payments ORDER BY id;`
	err = pr.db.SelectContext(ctx, &dest, query)
	if err != nil {
		return
	}
	if len(dest) == 0 {
		err = payment.ErrorNotFound
		return
	}
	return
}

func (pr *PaymentRepository) Update(ctx context.Context, id string, data payment.Entity) (err error) {
	sets, args := pr.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, id)
		query := fmt.Sprintf("UPDATE payments SET %s WHERE id = $%d RETURNING id;", strings.Join(sets, ","), len(args))
		if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = payment.ErrorNotFound
			}
		}
	}
	return
}

func (pr *PaymentRepository) prepareArgs(data payment.Entity) (sets []string, args []any) {
	if data.UserID != "" {
		sets = append(sets, "user_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, data.UserID)
	}
	if len(data.OrderID) != 0 {
		sets = append(sets, "order_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, pq.Array(data.OrderID))
	}
	if data.Amount != 0 {
		sets = append(sets, "amount = $"+strconv.Itoa(len(args)+1))
		args = append(args, data.Amount)
	}
	if data.Status != "" {
		sets = append(sets, "status = $"+strconv.Itoa(len(args)+1))
		args = append(args, data.Status)
	}
	return
}

func (pr *PaymentRepository) Search(ctx context.Context, filter, value string) (payments []payment.Entity, err error) {
	payments = []payment.Entity{}

	log.Printf("filter: %s, value: %s", filter, value)
	filter = pr.prepareFilter(filter)
	query := fmt.Sprintf("SELECT * FROM payments WHERE %s = $1;", filter)
	err = pr.db.SelectContext(ctx, &payments, query, value)
	if err != nil {
		return
	}

	if len(payments) == 0 {
		err = payment.ErrorNotFound
		return
	}
	return
}

func (pr *PaymentRepository) prepareFilter(filter string) string {
	switch filter {
	case "user_id":
		return "user_id"
	case "order_id":
		return "order_id"
	case "status":
		return "status"
	default:
		return ""
	}
}
