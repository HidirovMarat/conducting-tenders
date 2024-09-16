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
	TenderId    uuid.UUID             `db:"tender_id"`
	AuthorType  authorType.AuthorType `db:"author_t"`
	AuthorId    uuid.UUID             `db:"author_id"`
	Version     int                   `db:"version"`
	CreatedAt   time.Time             `db:"created_at"`
	Tag         uuid.UUID             `db:"tag" json:"-"`
}
