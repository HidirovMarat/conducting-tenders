package service

import (
	"conducting-tenders/internal/entity"
	"conducting-tenders/internal/entity/statusTender"
	"conducting-tenders/internal/repo"
	"conducting-tenders/internal/repo/repoerrs"
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type TenderService struct {
	tenderRepo                  repo.Tender
	employeeRepo                repo.Employee
	organizationResponsibleRepo repo.OrganizationResponsible
	organizationRepo            repo.Organization
}

func NewTenderService(tenderRepo repo.Tender, employeeRepo repo.Employee, organizationResponsibleRepo repo.OrganizationResponsible, organizationRepo repo.Organization) *TenderService {
	return &TenderService{
		tenderRepo: tenderRepo,
		employeeRepo: employeeRepo,
		organizationResponsibleRepo: organizationResponsibleRepo,
		organizationRepo: organizationRepo,
	}
}

func (b *TenderService) CreateTender(ctx context.Context, createTenderInput CreateTenderInput) (entity.Tender, error) {
	var tender entity.Tender

	tender.Id = uuid.New()
	tender.Name = createTenderInput.Name
	tender.Description = createTenderInput.Description
	tender.ServiceType = createTenderInput.ServiceType
	tender.Status = statusTender.StatusCreated

	_, err := b.organizationRepo.GetOrganizationById(ctx, createTenderInput.OrganizationId)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrOrganizationNotFind
		}
		return entity.Tender{}, err
	}
	tender.OrganizationId = createTenderInput.OrganizationId
	tender.Version = 1
	tender.CreatedAt = time.Now()
	tender.Tag = uuid.New()

	_, err = b.tenderRepo.CreateTender(ctx, tender)

	if err != nil {
		return entity.Tender{}, err
	}

	return tender, nil
}

func (b *TenderService) GetTendersByUsername(ctx context.Context, input GetTendersByUsernameInput) ([]entity.Tender, error) {
	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return []entity.Tender{}, ErrEmployeeNotFind
		}
		return []entity.Tender{}, err
	}

	organizationId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return []entity.Tender{}, ErrOrganizationNotFind
		}
		return []entity.Tender{}, err
	}

	if input.Limit == 0 {
		input.Limit = 5
	}

	tenders, err := b.tenderRepo.GetTendersByOrganizationId(ctx, organizationId, input.Limit, input.Offset)

	if err != nil {
		return []entity.Tender{}, err
	}

	return tenders, nil
}

func (b *TenderService) GetTenders(ctx context.Context, input GetTendersInput) ([]entity.Tender, error) {
	if input.Limit == 0 {
		input.Limit = 5
	}

	tenders, err := b.tenderRepo.GetTenders(ctx, input.ServiceType, input.Limit, input.Offset)
	if err != nil {
		return []entity.Tender{}, err
	}

	return tenders, nil
}

func (b *TenderService) GetTenderStatusById(ctx context.Context, input GetTenderStatusByIdInput) (statusTender.Status, error) {
	tender, err := b.tenderRepo.GetTenderById(ctx, input.TenderId)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return "", ErrTenderNotFind
		}
		return "", err
	}

	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return "", ErrEmployeeNotFind
		}
		return "", err
	}

	if tender.Status == statusTender.StatusPublished {
		return tender.Status, nil
	}

	orgaEmpId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return "", ErrOrganizationNotFind
		}
		return "", err
	}
	log.Println("to = ", tender.ServiceType)
	log.Println("org =  %s",orgaEmpId)
	if tender.OrganizationId != orgaEmpId {
		return "", ErrNotEnoughRights
	}

	return tender.Status, nil
}

func (b *TenderService) UpdateTenderStatusById(ctx context.Context, input UpdateTenderStatusByIdInput) (entity.Tender, error) {
	tender, err := b.tenderRepo.GetTenderById(ctx, input.TenderId)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrTenderNotFind
		}
		return entity.Tender{}, err
	}

	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrEmployeeNotFind
		}
		return entity.Tender{}, err
	}

	orgaEmpId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrOrganizationNotFind
		}
		return entity.Tender{}, err
	}

	if tender.OrganizationId != orgaEmpId {
		return entity.Tender{}, ErrNotEnoughRights
	}
	maxVersion, err := b.tenderRepo.GetTenderVersionMaxByTag(ctx, tender.Tag)
	if err != nil {
		return entity.Tender{}, err
	}
	tender.Status = input.Status
	tender.Version = maxVersion + 1
	tender.Id = uuid.New()
	tenderId, err := b.tenderRepo.CreateTender(ctx, tender)

	if err != nil {
		return entity.Tender{}, err
	}

	tender.Id = tenderId

	return tender, nil
}

func (b *TenderService) EditTenderById(ctx context.Context, input EditTenderByIdInput) (entity.Tender, error) {
	tender, err := b.tenderRepo.GetTenderById(ctx, input.TenderId)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrTenderNotFind
		}
		return entity.Tender{}, err
	}

	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrEmployeeNotFind
		}
		return entity.Tender{}, err
	}

	orgaEmpId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrOrganizationNotFind
		}
		return entity.Tender{}, err
	}

	if tender.OrganizationId != orgaEmpId {
		return entity.Tender{}, ErrNotEnoughRights
	}

	if input.Name == "" && input.Description == "" && input.ServiceType == "" {
		return entity.Tender{}, ErrInvalidRequestFormatOrParameters
	}

	if input.Name != "" {
		tender.Name = input.Name
	}

	if input.Description != "" {
		tender.Description = input.Description
	}

	if input.ServiceType != "" {
		tender.ServiceType = input.ServiceType
	}
	maxVersion, err := b.tenderRepo.GetTenderVersionMaxByTag(ctx, tender.Tag)
	if err != nil {
		return entity.Tender{}, err
	}
	tender.Version = maxVersion + 1
	tender.Id = uuid.New()
	tenderId, err := b.tenderRepo.CreateTender(ctx, tender)
	if err != nil {
		return entity.Tender{}, err
	}

	tender.Id = tenderId

	return tender, nil
}

func (b *TenderService) UpdateVersionTender(ctx context.Context, input UpdateVersionTenderInput) (entity.Tender, error) {
	tender, err := b.tenderRepo.GetTenderById(ctx, input.TenderId)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrTenderNotFind
		}
		return entity.Tender{}, err
	}

	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrEmployeeNotFind
		}
		return entity.Tender{}, err
	}

	orgaEmpId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrOrganizationNotFind
		}
		return entity.Tender{}, err
	}

	if tender.OrganizationId != orgaEmpId {
		return entity.Tender{}, ErrNotEnoughRights
	}

	tenderVer, err := b.tenderRepo.GetTenderByTagAndVersion(ctx, tender.Tag, input.Version )

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Tender{}, ErrTenderNotFind
		}
		return entity.Tender{}, err
	}
	maxVersion, err := b.tenderRepo.GetTenderVersionMaxByTag(ctx, tender.Tag)
	if err != nil {
		return entity.Tender{}, err
	}
	tenderVer.Version = maxVersion + 1 
	tenderVer.Id = uuid.New()
	tenderVerId, err := b.tenderRepo.CreateTender(ctx, tenderVer)

	if err != nil {
		return entity.Tender{}, err
	}

	tenderVer.Id = tenderVerId 

	return tenderVer, nil
}