package service

import (
	"conducting-tenders/internal/entity"
	"conducting-tenders/internal/repo"
	"context"
)

type OrganizationResponsibleService struct {
	organizationResponsibleRepo repo.OrganizationResponsible
}

func NewOrganizationResponsibleService(organizationResponsibleRepo repo.OrganizationResponsible) *OrganizationResponsibleService {
	return &OrganizationResponsibleService{
		organizationResponsibleRepo: organizationResponsibleRepo,
	}
}

func (b *OrganizationResponsibleService) GetOrganizationResponsibles(ctx context.Context, input GetOrganizationResponsiblesInput) ([]entity.OrganizationResponsible, error) {
	organizationResponsibles, err := b.organizationResponsibleRepo.GetOrganizationResponsibles(ctx, input.Limit)

	if err != nil {
		return []entity.OrganizationResponsible{}, err
	}

	return organizationResponsibles, nil
}
