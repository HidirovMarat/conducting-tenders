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

type OrganizationResponsibleRepo struct {
	*postgres.Postgres
}

func NewOrganizationResponsibleRepo(pg *postgres.Postgres) *OrganizationResponsibleRepo {
	return &OrganizationResponsibleRepo{pg}
}

func (r *OrganizationResponsibleRepo) GetOrganizationIdByEmployeeId(ctx context.Context, employeeId uuid.UUID) (uuid.UUID, error) {
	sql, args, _ := r.Builder.
		Select("organization_id").
		From("organization_responsible").
		Where("user_id  = ?", employeeId).
		ToSql()

	var organizationResponsible entity.OrganizationResponsible
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&organizationResponsible.Id,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.Nil, repoerrs.ErrNotFound
		}
		return uuid.Nil, fmt.Errorf("OrganizationResponsibleRepo.GetOrganizationIdByEmployeeId - r.Pool.QueryRow: %v", err)
	}

	return organizationResponsible.Id, nil
}

func(r *OrganizationResponsibleRepo) 	GetOrganizationResponsibles(ctx context.Context, limit int) ([]entity.OrganizationResponsible, error) {
	req := r.Builder.
		Select("*").
		From("organization_responsible")

	if limit > 0 {
		req.Limit(uint64(limit))
	}

	sql, args, _ := req.ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return []entity.OrganizationResponsible{}, fmt.Errorf("OrganizationResponsibleRepo.GetOrganizationResponsibles - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var organizationResponsibles []entity.OrganizationResponsible
	for rows.Next() {
		var organizationResponsible entity.OrganizationResponsible
		err := rows.Scan(
			&organizationResponsible.Id,
			&organizationResponsible.OrganizationId,
			&organizationResponsible.UserId,
		)
		if err != nil {
			return nil, fmt.Errorf("OrganizationResponsibleRepo.GetOrganizationResponsibles - rows.Scan: %v", err)
		}
		organizationResponsibles = append(organizationResponsibles, organizationResponsible)
	}

	return organizationResponsibles, nil
}
