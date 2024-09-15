package pgdb

import (
	"conducting-tenders/pkg/postgres"
	"context"
)

type PingRepo struct {
	*postgres.Postgres
}

func NewPingRepo(pg *postgres.Postgres) *PingRepo {
	return &PingRepo{pg}
}

func (r *PingRepo) ChechDb(ctx context.Context) error {
	return r.Pool.Ping(ctx)
}
