package user

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
	"time"
)

var (
	ErrorNotFound       = errors.New("user not found")
	ErrorInvalidSearch  = errors.New("invalid search parameters")
	ErrorInvalidName    = errors.New("invalid name")
	ErrorInvalidEmail   = errors.New("invalid email")
	ErrorInvalidRole    = errors.New("invalid role")
	ErrorInvalidAddress = errors.New("invalid address")
)

type Request struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Roles   string `json:"roles"`
}

type Response struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Address string    `json:"address"`
	RegDate time.Time `json:"reg_date"`
	Roles   string    `json:"roles"`
}

func ParseFromEntity(entity Entity) Response {
	return Response{
		ID:      entity.ID,
		Name:    entity.Name,
		Email:   entity.Email,
		Address: entity.Address,
		RegDate: entity.RegDate,
		Roles:   entity.Roles,
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
	if r.Name == "" {
		return ErrorInvalidName
	}
	if r.Address == "" {
		return ErrorInvalidAddress
	}
	if !isValidEmail(r.Email) {
		return ErrorInvalidEmail
	}
	if !isValidRole(r.Roles) {
		return ErrorInvalidRole
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func isValidRole(role string) bool {
	validRoles := map[string]bool{
		"admin":     true,
		"user":      true,
		"manager":   true,
		"developer": true,
	}
	return validRoles[role]
}

func IsValidFilter(filter string) bool {
	switch filter {
	case "name", "email":
		return true
	default:
		return false
	}
}
