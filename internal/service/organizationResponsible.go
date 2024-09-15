package service

import (
	"conducting-tenders/internal/entity"
	"conducting-tenders/internal/repo"
	"context"
)

type OrganizationService struct {
	organizationRepo repo.Organization
}

func NewOrganizationService(organizationRepo repo.Organization) *OrganizationService {
	return &OrganizationService{
		organizationRepo: organizationRepo,
	}
}

func (b *OrganizationService) GetOrganizations(ctx context.Context, input GetOrganizationsInput) ([]entity.Organization, error) {
	organizations, err := b.organizationRepo.GetOrganizations(ctx, input.Limit)

	if err != nil {
		return []entity.Organization{}, err
	}

	return organizations, nil
}