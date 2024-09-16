package pgdb

import (
	"conducting-tenders/internal/entity"
	serviceType "conducting-tenders/internal/entity/service-type"
	"conducting-tenders/internal/repo/repoerrs"
	"conducting-tenders/pkg/postgres"
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
)

type TenderRepo struct {
	*postgres.Postgres
}

func NewTenderRepo(pg *postgres.Postgres) *TenderRepo {
	return &TenderRepo{pg}
}

func (r *TenderRepo) CreateTender(ctx context.Context, tender entity.Tender) (uuid.UUID, error) {
	sql, args, _ := r.Builder.
		Insert("tenders").
		Columns("id", "name", "description", "type", "status", "organization_id", "version", "created_at", "tag").
		Values(tender.Id, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationId, tender.Version, tender.CreatedAt, tender.Tag).
		Suffix("RETURNING id").
		ToSql()

	var id uuid.UUID
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("TenderRepo.CreateTender - r.Pool.QueryRow: %v", err)
	}

	return id, nil
}

func (r *TenderRepo) GetTenderById(ctx context.Context, tenderId uuid.UUID) (entity.Tender, error) {
	sql, args, _ := r.Builder.
		Select("id", "name", "description", "type", "status", "organization_id", "version", "created_at", "tag").
		From("tenders").
		Where("id = ?", tenderId).
		ToSql()

	var tender entity.Tender
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&tender.Id,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
		&tender.Status,
		&tender.OrganizationId,
		&tender.Version,
		&tender.CreatedAt,
		&tender.Tag,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Tender{}, repoerrs.ErrNotFound
		}
		return entity.Tender{}, fmt.Errorf("TenderRepo.GetTenderById - r.Pool.QueryRow: %v", err)
	}

	return tender, nil
}

func (r *TenderRepo) GetTendersByOrganizationId(ctx context.Context, organizationId uuid.UUID, limit int, offset int) ([]entity.Tender, error) {
	req := r.Builder.
		Select("id", "name", "description", "type", "status", "organization_id", "version", "created_at").
		From("tenders").
		Where("organization_id = ?", organizationId).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	sql, args, _ := req.ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("TenderRepo.GetTendersByOrganizationId - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var tenders []entity.Tender
	for rows.Next() {
		var tender entity.Tender
		err := rows.Scan(
			&tender.Id,
			&tender.Name,
			&tender.Description,
			&tender.ServiceType,
			&tender.Status,
			&tender.OrganizationId,
			&tender.Version,
			&tender.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("TenderRepo.GetTendersByOrganizationId - rows.Scan: %v", err)
		}
		tenders = append(tenders, tender)
	}

	return tenders, nil
}

func (r *TenderRepo) GetTenders(ctx context.Context, serviceType []serviceType.ServiceType, limit int, offset int) ([]entity.Tender, error) {
	req := r.Builder.
		Select("id", "name", "description", "type", "status", "organization_id", "version", "created_at").
		From("tenders").
		Where(sq.Eq{"type": serviceType}).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	sql, args, _ := req.ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("TenderRepo.GetTenders - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var tenders []entity.Tender
	for rows.Next() {
		var tender entity.Tender
		err := rows.Scan(
			&tender.Id,
			&tender.Name,
			&tender.Description,
			&tender.ServiceType,
			&tender.Status,
			&tender.OrganizationId,
			&tender.Version,
			&tender.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("TenderRepo.GetTenders - rows.Scan: %v", err)
		}
		tenders = append(tenders, tender)
	}

	return tenders, nil
}

func (r *TenderRepo) GetTenderByTagAndVersion(ctx context.Context, tag uuid.UUID, version int) (entity.Tender, error) {
	sql, args, _ := r.Builder.
		Select("id", "name", "description", "type", "status", "organization_id", "version", "created_at", "tag").
		From("tenders").
		Where("tag = ? and version = ?", tag, version).
		ToSql()

	var tender entity.Tender
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&tender.Id,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
		&tender.Status,
		&tender.OrganizationId,
		&tender.Version,
		&tender.CreatedAt,
		&tender.Tag,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Tender{}, repoerrs.ErrNotFound
		}
		return entity.Tender{}, fmt.Errorf("TenderRepo.GetTenderByTagAndVersion - r.Pool.QueryRow: %v", err)
	}

	return tender, nil
}

func (r *TenderRepo) GetTenderVersionMaxByTag(ctx context.Context, tag uuid.UUID) (int, error) {
	sql, args, _ := r.Builder.
		Select("MAX(version) as max").
		From("tenders").
		Where("tag = ?", tag).
		ToSql()

	var max int
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&max,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, repoerrs.ErrNotFound
		}
		return 0, fmt.Errorf("TenderRepo.GetTenderVersionMaxByTag - r.Pool.QueryRow: %v", err)
	}

	return max, nil
}
