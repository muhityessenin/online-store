package order

type Request struct {
	UserID    string   `db:"user_id" bson:"user_id"`
	ProductID []string `db:"product_id" bson:"product_id"`
	Pricing   float64  `db:"pricing" bson:"pricing"`
	Status    string   `db:"status" bson:"status"`
}
