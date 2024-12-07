package product

import "time"

type Entity struct {
	ID          string    `db:"id" bson:"_id"`
	Title       string    `db:"title" bson:"title"`
	Description string    `db:"description" bson:"description"`
	Price       float64   `db:"price" bson:"price"`
	Category    string    `db:"category" bson:"category"`
	Quantity    int       `db:"quantity" bson:"quantity"`
	CreatedAt   time.Time `db:"created_at" bson:"created_at"`
}
