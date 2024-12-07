package user

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID      uuid.UUID `db:"id" bson:"_id"`
	Name    string    `db:"name" bson:"name"`
	Email   string    `db:"email" bson:"email"`
	Address string    `db:"address" bson:"address"`
	RegDate time.Time `db:"reg_date" bson:"reg_date"`
	Roles   string    `db:"roles" bson:"roles"`
}
