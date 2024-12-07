package order

import (
	"github.com/lib/pq"
	"time"
)

type Entity struct {
	ID        string         `db:"id" bson:"_id"`
	UserID    string         `db:"user_id" bson:"user_id"`
	ProductID pq.StringArray `db:"product_id" bson:"product_id"`
	Pricing   float64        `db:"pricing" bson:"pricing"`
	Status    string         `db:"status" bson:"status"`
	CreatedAt time.Time      `db:"created_at" bson:"created_at"`
}
