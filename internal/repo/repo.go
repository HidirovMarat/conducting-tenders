package repo

import (
	"conducting-tenders/internal/entity"
	serviceType "conducting-tenders/internal/entity/service-type"
	"conducting-tenders/internal/repo/pgdb"
	"conducting-tenders/pkg/postgres"
	"context"

	"github.com/google/uuid"
)

type Ping interface {
	ChechDb(ctx context.Context) error
}

type Bid interface {
	CreateBid(ctx context.Context, bid entity.Bid) (uuid.UUID, error)
	GetBidsByAuthorId(ctx context.Context, authorId uuid.UUID, limit int, offset int) ([]entity.Bid, error)
	GetBidsByTenderIdAndAuthorId(ctx context.Context, tenderId uuid.UUID, authorId uuid.UUID, limit int, offset int) ([]entity.Bid, error)
	GetBidById(ctx context.Context, bidId uuid.UUID) (entity.Bid, error)
	GetBidByTagAndVersion(ctx context.Context, tag uuid.UUID, version int) (entity.Bid, error)
}

type Tender interface {
	CreateTender(ctx context.Context, tender entity.Tender) (uuid.UUID, error)
	GetTenders(ctx context.Context, serviceType []serviceType.ServiceType, limit int, offset int) ([]entity.Tender, error)
	GetTendersByOrganizationId(ctx context.Context, organizationId uuid.UUID, limit int, offset int) ([]entity.Tender, error)
	GetTenderById(ctx context.Context, tenderId uuid.UUID) (entity.Tender, error)
	GetTenderByTagAndVersion(ctx context.Context, tag uuid.UUID, version int) (entity.Tender, error)
}

type Employee interface {
	GetEmployeeByUsername(ctx context.Context, username string) (entity.Employee, error)
	GetEmployeeById(ctx context.Context, employeeId uuid.UUID) (entity.Employee, error)
	GetEmployees(ctx context.Context, limit int) ([]entity.Employee, error)
}

type Organization interface {
	GetOrganizationById(ctx context.Context, organizationId uuid.UUID) (entity.Organization, error)
	GetOrganizations(ctx context.Context, limit int) ([]entity.Organization, error)
}

type OrganizationResponsible interface {
	GetOrganizationIdByEmployeeId(ctx context.Context, employeeId uuid.UUID) (uuid.UUID, error)
	GetOrganizationResponsibles(ctx context.Context, limit int) ([]entity.OrganizationResponsible, error)
}

type Repositories struct {
	Ping
	Bid
	Employee
	Tender
	Organization
	OrganizationResponsible
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Bid:      pgdb.NewBidRepo(pg),
		Tender:   pgdb.NewTenderRepo(pg),
		Employee: pgdb.NewEmployeeRepo(pg),
		Ping:     pgdb.NewPingRepo(pg),
		Organization: pgdb.NewOrganizationRepo(pg),
		OrganizationResponsible: pgdb.NewOrganizationResponsibleRepo(pg),
	}
}
