package service

import (
	"conducting-tenders/internal/repo"
	"context"
)

type PingService struct {
	pingRepo repo.Ping
}

func NewPingService(pingRepo repo.Ping) *PingService {
	return &PingService{
		pingRepo: pingRepo,
	}
}

func (s *PingService) Ping(ctx context.Context) error {
	return s.pingRepo.ChechDb(ctx)
}
