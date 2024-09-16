package entity

import (
	serviceType "conducting-tenders/internal/entity/service-type"
	"conducting-tenders/internal/entity/statusTender"
	"time"

	"github.com/google/uuid"
)

type Tender struct {
	Id             uuid.UUID               `db:"id"`
	Name           string                  `db:"name"`
	Description    string                  `db:"description"`
	ServiceType    serviceType.ServiceType `db:"type"`
	Status         statusTender.Status     `db:"status"`
	OrganizationId uuid.UUID               `db:"organization_id"`
	Version        int                     `db:"version"`
	CreatedAt      time.Time               `db:"created_at"`
	Tag            uuid.UUID               `db:"tag" json:"-"`
}
