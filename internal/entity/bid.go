package entity

import (
	authorType "conducting-tenders/internal/entity/author-type"
	"conducting-tenders/internal/entity/statusBid"
	"time"

	"github.com/google/uuid"
)

type Bid struct {
	Id          uuid.UUID             `db:"id"`
	Name        string                `db:"name"`
	Description string                `db:"description"`
	Status      statusBid.Status      `db:"statuc"`
	TenderId    uuid.UUID             `db:"tenderId"`
	AuthorType  authorType.AuthorType `db:"authorType"`
	AuthorId    uuid.UUID             `db:"authorId"`
	Version     int                   `db:"version"`
	CreatedAt   time.Time             `db:"createdAt"`
	Tag         uuid.UUID             `db:"tag"`
}
