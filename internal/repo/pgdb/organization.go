package pgdb

import (
	"conducting-tenders/internal/entity"
	"conducting-tenders/internal/repo/repoerrs"
	"conducting-tenders/pkg/postgres"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type OrganizationRepo struct {
	*postgres.Postgres
}

func NewOrganizationRepo(pg *postgres.Postgres) *OrganizationRepo {
	return &OrganizationRepo{pg}
}

/*
	Id               uuid.UUID `db:"id"`
	Name             string    `db:"name"`
	Description      string    `db:"description"`
	OrganizationType string    `db:"organization_type"` //TODO
	Created_at       time.Time `db:"created_at"`
	Updated_at       time.Time `db:"updated_at"`
*/

func (r *OrganizationRepo) GetOrganizationById(ctx context.Context, organizationId uuid.UUID) (entity.Organization, error) {
	sql, args, _ := r.Builder.
		Select("id", "name", "description", "type", "created_at", "updated_at").
		From("organization").
		Where("id  = ?", organizationId).
		ToSql()

	var organization entity.Organization
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&organization.Id,
		&organization.Name,
		&organization.Description,
		&organization.OrganizationType,
		&organization.Created_at,
		&organization.Updated_at,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Organization{}, repoerrs.ErrNotFound
		}
		return entity.Organization{}, fmt.Errorf("OrganizationRepo.GetOrganizationById - r.Pool.QueryRow: %v", err)
	}

	return entity.Organization{}, nil
}

func (r *OrganizationRepo) GetOrganizations(ctx context.Context, limit int) ([]entity.Organization, error) {
	req := r.Builder.
		Select("id", "name", "description").
		From("organization")

	if limit > 0 {
		req.Limit(uint64(limit))
	}

	sql, args, _ := req.ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return []entity.Organization{}, fmt.Errorf("OrganizationRepo.GetOrganizations - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var organizations []entity.Organization
	for rows.Next() {
		var organization entity.Organization
		err := rows.Scan(
			&organization.Id,
			&organization.Name,
			&organization.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("OrganizationRepo.GetOrganizations - rows.Scan: %v", err)
		}
		organizations = append(organizations, organization)
	}

	return organizations, nil
}
