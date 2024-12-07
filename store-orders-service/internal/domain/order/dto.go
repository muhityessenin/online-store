package order

import (
	"errors"
	"time"
)

var (
	ErrorNotFound         = errors.New("payment not found")
	ErrorInvalidStatus    = errors.New("invalid status")
	ErrorInvalidPrice     = errors.New("invalid price")
	ErrorInvalidSearch    = errors.New("invalid search filter")
	ErrorInvalidUserID    = errors.New("invalid user id")
	ErrorInvalidProductID = errors.New("invalid product id")
)

type Request struct {
	UserID    string   `db:"user_id" bson:"user_id"`
	ProductID []string `db:"product_id" bson:"product_id"`
	Pricing   float64  `db:"pricing" bson:"pricing"`
	Status    string   `db:"status" bson:"status"`
}

type Response struct {
	ID        string    `json:"id"`
	UserID    string    `db:"user_id" bson:"user_id"`
	ProductID []string  `db:"product_id" bson:"product_id"`
	Pricing   float64   `db:"pricing" bson:"pricing"`
	Status    string    `db:"status" bson:"status"`
	CreatedAt time.Time `db:"created_at" bson:"created_at"`
}

func ParseFromEntity(entity Entity) Response {
	return Response{
		ID:        entity.ID,
		UserID:    entity.UserID,
		ProductID: entity.ProductID,
		Pricing:   entity.Pricing,
		Status:    entity.Status,
		CreatedAt: entity.CreatedAt,
	}
}

func ParseFromEntities(data []Entity) (res []Response) {
	res = make([]Response, 0)
	for _, entity := range data {
		res = append(res, ParseFromEntity(entity))
	}
	return
}

func (r *Request) Validate() error {
	if r.UserID == "" {
		return ErrorInvalidUserID
	}
	if len(r.ProductID) == 0 {
		return ErrorInvalidProductID

	}
	if r.Pricing <= 0 {
		return ErrorInvalidPrice
	}
	if !isValidStatus(r.Status) {
		return ErrorInvalidStatus
	}
	return nil
}

func IsValidFilter(filter string) bool {
	return filter == "user_id" || filter == "status"
}

func isValidStatus(status string) bool {
	return status == "new" || status == "in_progress" || status == "done"
}
