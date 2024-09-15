package entity

import (
	"github.com/google/uuid"
)

type OrganizationResponsible struct {
	Id             uuid.UUID `db:"id"`
	OrganizationId uuid.UUID `db:"organization_id"`
	UserId         uuid.UUID `db:"user_id"`
}
