package entity

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	Id               uuid.UUID `db:"id"`
	Name             string    `db:"name"`
	Description      string    `db:"description"`
	OrganizationType string    `db:"type"` //TODO
	Created_at       time.Time `db:"created_at"`
	Updated_at       time.Time `db:"updated_at"`
}
