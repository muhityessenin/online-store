package product

import (
	"errors"
	"time"
)

var (
	ErrorNotFound           = errors.New("payment not found")
	ErrorInvalidTitle       = errors.New("invalid title")
	ErrorInvalidDescription = errors.New("invalid description")
	ErrorInvalidDate        = errors.New("invalid date format")
	ErrorInvalidStatus      = errors.New("invalid status")
	ErrorInvalidPrice       = errors.New("invalid price")
	ErrorInvalidCategory    = errors.New("invalid category")
	ErrorInvalidQuantity    = errors.New("invalid quantity")
	ErrorInvalidSearch      = errors.New("invalid search filter")
)

type Request struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Quantity    int     `json:"quantity"`
}

type Response struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

func ParseFromEntity(entity Entity) Response {
	return Response{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Price:       entity.Price,
		Category:    entity.Category,
		Quantity:    entity.Quantity,
		CreatedAt:   entity.CreatedAt,
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
	if r.Title == "" && len(r.Title) < 200 {
		return ErrorInvalidTitle
	}
	if r.Description == "" {
		return ErrorInvalidDescription
	}
	if r.Price <= 0 {
		return ErrorInvalidPrice
	}
	if r.Category == "" {
		return ErrorInvalidCategory
	}
	if r.Quantity <= 0 {
		return ErrorInvalidQuantity
	}
	return nil
}

func ParseDate(date string) (data time.Time) {
	data, _ = time.Parse("2006-01-02", date)
	return
}

func IsValidFilter(filter string) bool {
	return filter == "title" || filter == "category"
}
