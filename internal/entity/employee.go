package entity

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	Id         uuid.UUID `db:"id"`
	Username   string    `db:"username"`
	First_name string    `db:"first_name"`
	Last_name  string    `db:"last_name"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}
