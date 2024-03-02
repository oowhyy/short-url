package api

import (
	"github.com/oowhyy/short-url/internal/service"
	"github.com/rs/zerolog"
)

type Server struct {
	listenAddr      string
	logger          zerolog.Logger
	shortUrlService service.ShortUrlService
}

func NewServer(config *Config, logger zerolog.Logger, shortUrlService service.ShortUrlService) *Server {
	return &Server{
		listenAddr:      config.ListenAddr,
		logger:          logger,
		shortUrlService: shortUrlService,
	}
}
