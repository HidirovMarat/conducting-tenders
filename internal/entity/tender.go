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
	ServiceType    serviceType.ServiceType `db:"serviceType"`
	Status         statusTender.Status     `db:"status"`
	OrganizationId uuid.UUID               `db:"organizationId"`
	Version        int                     `db:"version"`
	CreatedAt      time.Time               `db:"createdAt"`
	Tag            uuid.UUID               `db:"tag"`
}
