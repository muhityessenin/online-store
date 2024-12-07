package payment

import (
	"time"
)

type Entity struct {
	ID        string    `db:"id" bson:"_id"`
	UserID    string    `db:"user_id" bson:"user_id"`
	OrderID   string    `db:"order_id" bson:"order_id"`
	Amount    float64   `db:"amount" bson:"amount"`
	Status    string    `db:"status" bson:"status"`
	CreatedAt time.Time `db:"created_at" bson:"created_at"`
}
