package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
	"users-service/internal/domain/user"
	interfaces "users-service/internal/repository/interface"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(ctx context.Context, data user.Entity) (id string, err error) {
	query := `INSERT INTO users (name, address, email, roles) VALUES ($1,$2,$3,$4) RETURNING id;`
	args := []any{
		data.Name,
		data.Address,
		data.Email,
		data.Roles,
	}
	if err = ur.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if err != nil {
			err = user.ErrorNotFound
		}
	}
	return
}

func (ur *UserRepository) Delete(ctx context.Context, id string) (err error) {
	query := `DELETE FROM users WHERE id = $1 RETURNING id;`
	args := []any{id}

	if err = ur.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if err != nil {
			err = user.ErrorNotFound
		}
	}
	return
}

func (ur *UserRepository) Get(ctx context.Context, id string) (dest user.Entity, err error) {
	query := `SELECT * FROM users WHERE id = $1;`
	args := []any{id}
	err = ur.db.GetContext(ctx, &dest, query, args...)
	if err != nil {
		if err != nil {
			err = user.ErrorNotFound
		}
	}
	return
}

func (ur *UserRepository) List(ctx context.Context) (users []user.Entity, err error) {
	query := `SELECT * FROM users;`
	err = ur.db.SelectContext(ctx, &users, query)
	if err != nil {
		if err != nil {
			err = user.ErrorNotFound
		}
	}
	return
}

func (ur *UserRepository) Update(ctx context.Context, id string, data user.Entity) (err error) {
	sets, args := ur.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, id)
	}
	query := fmt.Sprintf(`UPDATE users SET %s WHERE id = $%d;`, strings.Join(sets, ","), len(args))
	_, err = ur.db.ExecContext(ctx, query, args...)
	if err != nil {
		if err != nil {
			err = user.ErrorNotFound
		}
	}
	return
}

func (ur *UserRepository) prepareArgs(data user.Entity) (sets []string, args []any) {
	if data.Name != "" {
		args = append(args, data.Name)
		sets = append(sets, fmt.Sprintf("name = $%d", len(args)))
	}
	if data.Email != "" {
		args = append(args, data.Email)
		sets = append(sets, fmt.Sprintf("email = $%d", len(args)))
	}
	if data.Roles != "" {
		args = append(args, data.Roles)
		sets = append(sets, fmt.Sprintf("roles = $%d", len(args)))
	}
	if data.Address != "" {
		args = append(args, data.Address)
		sets = append(sets, fmt.Sprintf("address = $%d", len(args)))
	}
	return
}

func (ur *UserRepository) Search(ctx context.Context, filter, value string) (users []user.Entity, err error) {
	users = []user.Entity{}

	log.Printf("filter: %s, value: %s", filter, value)
	filter = ur.prepareFilter(filter)

	query := fmt.Sprintf("SELECT * FROM users WHERE %s = $1;", filter)
	err = ur.db.SelectContext(ctx, &users, query, value)
	if err != nil {
		return
	}

	if len(users) == 0 {
		err = user.ErrorNotFound
		return
	}
	return
}

func (ur *UserRepository) prepareFilter(filter string) string {
	switch filter {
	case "name":
		return "name"
	case "email":
		return "email"
	default:
		return ""
	}
}
