package service

import (
	"conducting-tenders/internal/entity"
	authorType "conducting-tenders/internal/entity/author-type"
	serviceType "conducting-tenders/internal/entity/service-type"
	"conducting-tenders/internal/entity/statusBid"
	"conducting-tenders/internal/entity/statusTender"
	"conducting-tenders/internal/repo"
	"context"

	"github.com/google/uuid"
)

type Ping interface {
	Ping(ctx context.Context) error
}

type CreateBidInput struct {
	Name        string                `json:"name" validate:"required,min=1,max=100"`
	Description string                `json:"description" validate:"required,min=1,max=500"`
	TenderId    uuid.UUID             `json:"tenderId" validate:"required, uuid"`
	AuthorType  authorType.AuthorType `json:"authorType" validate:"required, oneof=Organization User"`
	AuthorId    uuid.UUID             `json:"authorId" validate:"required, uuid"`
}

type GetBidsByUsernameInput struct {
	Limit    int    `query:"limit"`
	Offset   int    `query:"offset"`
	Username string `query:"username" validate:"required, min=1,max=100"`
}

type GetBidsByTenderIdInput struct {
	Limit    int       `query:"limit"`
	Offset   int       `query:"offset"`
	Username string    `query:"username" validate:"required, min=1,max=100"`
	TenderId uuid.UUID `param:"tenderId" validate:"required, uuid"`
}

type GetBidStatusByIdInput struct {
	BidId    uuid.UUID `param:"bidId" validate:"required, uuid"`
	Username string    `query:"username" validate:"required, min=1, max=100"`
}

type UpdateBidStatusByIdInput struct {
	BidId    uuid.UUID        `param:"bidId" validate:"required, uuid"`
	Username string           `query:"username" validate:"required, min=1,max=100"`
	Status   statusBid.Status `query:"status" validate:"required, oneof=Created Published Canceled"`
}

type EditBidByIdInput struct {
	BidId       uuid.UUID `param:"bidId" validate:"required, uuid"`
	Username    string    `query:"username" validate:"required, min=1, max=100"`
	Name        string    `json:"name" validate:"omitempty, min=1, max=100"`
	Description string    `json:"description" validate:"omitempty, min=1, max=500"`
}

type Bid interface {
	CreateBid(ctx context.Context, createBidInput CreateBidInput) (entity.Bid, error)
	GetBidsByUsername(ctx context.Context, input GetBidsByUsernameInput) ([]entity.Bid, error)
	GetBidsByTenderId(ctx context.Context, input GetBidsByTenderIdInput) ([]entity.Bid, error)
	GetBidStatusById(ctx context.Context, input GetBidStatusByIdInput) (statusBid.Status, error)
	UpdateBidStatusById(ctx context.Context, input UpdateBidStatusByIdInput) (entity.Bid, error)
	EditBidById(ctx context.Context, input EditBidByIdInput) (entity.Bid, error)
}

type CreateTenderInput struct {
	Name            string                  `json:"name" validate:"required,min=1,max=100"`
	Description     string                  `json:"description" validate:"required,min=1,max=500"`
	ServiceType     serviceType.ServiceType `json:"serviceType" validate:"required,oneof=Construction Delivery Manufacture"`
	OrganizationId  uuid.UUID               `json:"organizationId" validate:"required,uuid"`
	CreatorUsername string                  `json:"creatorUsername" validate:"required,min=1,max=100"`
}

type GetTendersInput struct {
	Limit       int                       `query:"limit"`
	Offset      int                       `query:"offset"`
	ServiceType []serviceType.ServiceType `query:"serviceType" validate:"dive,required,oneof=Construction Delivery Manufacture"`
}

type GetTendersByUsernameInput struct {
	Limit    int    `query:"limit"`
	Offset   int    `query:"offset"`
	Username string `query:"username" validate:"required, min=1, max=100"`
}

type GetTenderStatusByIdInput struct {
	TenderId uuid.UUID `param:"tenderId" validate:"required, uuid"`
	Username string    `query:"username" validate:"required, min=1, max=100"`
}

type UpdateTenderStatusByIdInput struct {
	TenderId uuid.UUID           `param:"tenderId" validate:"required, uuid"`
	Username string              `query:"username" validate:"required, min=1, max=100"`
	Status   statusTender.Status `query:"status" validate:"required, min=1, max=100"`
}

type EditTenderByIdInput struct {
	TenderId    uuid.UUID               `param:"tenderId" validate:"required, uuid"`
	Username    string                  `query:"username" validate:"required, min=1, max=100"`
	Name        string                  `json:"name" validate:"omitempty, min=1, max=100"`
	Description string                  `json:"description" validate:"omitempty, min=1, max=500"`
	ServiceType serviceType.ServiceType `json:"serviceType" validate:"omitempty, oneof=Construction Delivery Manufacture"`
}

type Tender interface {
	CreateTender(ctx context.Context, input CreateTenderInput) (entity.Tender, error)
	GetTenders(ctx context.Context, input GetTendersInput) ([]entity.Tender, error)
	GetTendersByUsername(ctx context.Context, input GetTendersByUsernameInput) ([]entity.Tender, error)
	GetTenderStatusById(ctx context.Context, input GetTenderStatusByIdInput) (statusTender.Status, error)
	UpdateTenderStatusById(ctx context.Context, input UpdateTenderStatusByIdInput) (entity.Tender, error)
	EditTenderById(ctx context.Context, input EditTenderByIdInput) (entity.Tender, error)
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

type GetEmployeesInput struct {
	Limit int `query:"limit"`
}

type Employee interface {
	GetEmployees(ctx context.Context, input GetEmployeesInput) ([]entity.Employee, error)
}

type GetOrganizationsInput struct {
	Limit int `query:"limit"`
}

type Organization interface {
	GetOrganizations(ctx context.Context, input GetOrganizationsInput) ([]entity.Organization, error)
}

type GetOrganizationResponsiblesInput struct {
	Limit int `query:"limit"`
}

type OrganizationResponsible interface {
	GetOrganizationResponsibles(ctx context.Context, input GetOrganizationResponsiblesInput) ([]entity.OrganizationResponsible, error)
}

type Services struct {
	Bid                     Bid
	Tender                  Tender
	Ping                    Ping
	Employee                Employee
	Organization            Organization
	OrganizationResponsible OrganizationResponsible
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Bid:                     NewBidService(deps.Repos, deps.Repos, deps.Repos, deps.Repos, deps.Repos),
		Tender:                  NewTenderService(deps.Repos, deps.Repos, deps.Repos, deps.Repos),
		Ping:                    NewPingService(deps.Repos),
		Employee:                NewEmployeeService(deps.Repos.Employee),
		Organization:            NewOrganizationService(deps.Repos),
		OrganizationResponsible: NewOrganizationResponsibleService(deps.Repos),
	}
}
