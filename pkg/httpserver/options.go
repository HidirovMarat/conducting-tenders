package httpserver

import (
	"time"
)

type Option func(*Server)

func ServerAddress(serverAddress string) Option {
	return func(s *Server) {
		s.Server.Addr = serverAddress
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Server.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Server.WriteTimeout = timeout
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
