package pgdb

import (
	"conducting-tenders/internal/entity"
	"conducting-tenders/internal/repo/repoerrs"
	"conducting-tenders/pkg/postgres"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
)

type BidRepo struct {
	*postgres.Postgres
}

func NewBidRepo(pg *postgres.Postgres) *BidRepo {
	return &BidRepo{pg}
}

func (r *BidRepo) CreateBid(ctx context.Context, bid entity.Bid) (uuid.UUID, error) {
	sql, args, _ := r.Builder.
		Insert("bids").
		Columns("id", "name", "description", "status", "tender_id", "author_t", "author_id", "version", "created_at", "tag").
		Values(bid.Id, bid.Name, bid.Description, bid.Status, bid.TenderId, bid.AuthorType, bid.AuthorId, bid.Version, bid.CreatedAt, bid.Tag).
		Suffix("RETURNING id").
		ToSql()

	var id uuid.UUID
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(&id)

	if err != nil {
		return uuid.Nil, fmt.Errorf("BidRepo.CreateBid - r.Pool.QueryRow: %v", err)
	}

	return id, nil
}

func (r *BidRepo) GetBidsByAuthorId(ctx context.Context, authorId uuid.UUID, limit int, offset int) ([]entity.Bid, error) {
	req := r.Builder.
		Select("id", "name", "description", "status", "tender_id", "author_t", "author_id", "version", "created_at").
		From("bids").
		Where("author_id = ?", authorId)

	if limit > 0 {
		req.Limit(uint64(limit))
	}

	if offset > 0 {
		req.Offset(uint64(offset))
	}
	sql, args, _ := req.ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("BidRepo.GetBidsByAuthorId - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var bids []entity.Bid
	for rows.Next() {
		var bid entity.Bid
		err := rows.Scan(
			&bid.Id,
			&bid.Name,
			&bid.Description,
			&bid.Status,
			&bid.TenderId,
			&bid.AuthorType,
			&bid.AuthorId,
			&bid.Version,
			&bid.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("BidRepo.GetBidsByAuthorId - rows.Scan: %v", err)
		}
		bids = append(bids, bid)
	}

	return bids, nil
}

func (r *BidRepo) GetBidsByTenderIdAndAuthorId(ctx context.Context, tenderId uuid.UUID, authorId uuid.UUID, limit int, offset int) ([]entity.Bid, error) {
	req := r.Builder.
		Select("id", "name", "description", "status", "tender_id", "author_t", "author_id", "version", "created_at").
		From("bids").
		Where("author_id = ? and tender_id = ?", authorId, tenderId)

	if limit > 0 {
		req.Limit(uint64(limit))
	}

	if offset > 0 {
		req.Offset(uint64(offset))
	}
	sql, args, _ := req.ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("BidRepo.GetBidsByTenderIdAndAuthorId - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var bids []entity.Bid
	for rows.Next() {
		var bid entity.Bid
		err := rows.Scan(
			&bid.Id,
			&bid.Name,
			&bid.Description,
			&bid.Status,
			&bid.TenderId,
			&bid.AuthorType,
			&bid.AuthorId,
			&bid.Version,
			&bid.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("BidRepo.GetBidsByTenderIdAndAuthorId - rows.Scan: %v", err)
		}
		bids = append(bids, bid)
	}

	return bids, nil
}

func (r *BidRepo) GetBidById(ctx context.Context, bidId uuid.UUID) (entity.Bid, error) {
	sql, args, _ := r.Builder.
		Select("id", "name", "description", "status", "tender_id", "author_t", "author_id", "version", "created_at", "tag").
		From("bids").
		Where("id = ?", bidId).
		ToSql()

	var bid entity.Bid
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&bid.Id,
		&bid.Name,
		&bid.Description,
		&bid.Status,
		&bid.TenderId,
		&bid.AuthorType,
		&bid.AuthorId,
		&bid.Version,
		&bid.CreatedAt,
		&bid.Tag,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Bid{}, repoerrs.ErrNotFound
		}
		return entity.Bid{}, fmt.Errorf("BidRepo.GetBidById - r.Pool.QueryRow: %v", err)
	}

	return bid, nil
}

func (r *BidRepo) GetBidByTagAndVersion(ctx context.Context, tag uuid.UUID, version int) (entity.Bid, error) {
	sql, args, _ := r.Builder.
		Select("id", "name", "description", "status", "tender_id", "author_t", "author_id", "version", "created_at", "tag").
		From("bids").
		Where("tag = ? and version = ?", tag, version).
		ToSql()

	var bid entity.Bid
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&bid.Id,
		&bid.Name,
		&bid.Description,
		&bid.Status,
		&bid.TenderId,
		&bid.AuthorType,
		&bid.AuthorId,
		&bid.Version,
		&bid.CreatedAt,
		&bid.Tag,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Bid{}, repoerrs.ErrNotFound
		}
		return entity.Bid{}, fmt.Errorf("BidRepo.GetBidByTagAndVersion - r.Pool.QueryRow: %v", err)
	}

	return bid, nil
}
