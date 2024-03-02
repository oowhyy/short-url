package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/oowhyy/short-url/internal/service"
	"github.com/oowhyy/short-url/pkg/shorturlpb"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server struct {
	listenAddr      string
	logger          zerolog.Logger
	shortUrlService service.ShortUrlService

	shorturlpb.UnimplementedShortUrlServer
}

func NewServer(config *Config, logger zerolog.Logger, shortUrlService service.ShortUrlService) *Server {
	return &Server{
		listenAddr:      config.ListenAddr,
		logger:          logger,
		shortUrlService: shortUrlService,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen tcp port: %w", err)
	}
	server := grpc.NewServer()
	shorturlpb.RegisterShortUrlServer(server, s)
	reflection.Register(server)
	s.logger.Info().Str("listenAddr", s.listenAddr).Msg("running grpc server")
	return server.Serve(lis)
}

func (s *Server) Shorten(ctx context.Context, req *shorturlpb.PostUrlRequest) (*shorturlpb.PostUrlResponse, error) {
	someString := req.GetOgUrl()
	res, err := s.shortUrlService.Shorten(someString)
	if err != nil {
		var serviceErr *service.Error
		ok := errors.As(err, &serviceErr)
		if !ok {
			return &shorturlpb.PostUrlResponse{}, status.Errorf(codes.Internal, "unknown error: %s", err)
		}
		if serviceErr.Reason == service.ReasonInvalidReq {
			return &shorturlpb.PostUrlResponse{}, status.Errorf(codes.InvalidArgument, "%s error: %s", serviceErr.Reason, serviceErr.Err)
		}
		return &shorturlpb.PostUrlResponse{}, status.Errorf(codes.Internal, "%s error: %s", serviceErr.Reason, serviceErr.Err)
	}
	return &shorturlpb.PostUrlResponse{ShortLink: res}, nil
}

func (s *Server) Reverse(ctx context.Context, req *shorturlpb.GetUrlRequest) (*shorturlpb.GetUrlResponse, error) {
	short := req.GetShortLink()
	res, err := s.shortUrlService.Reverse(short)
	if err != nil {
		var serviceErr *service.Error
		ok := errors.As(err, &serviceErr)
		if !ok {
			return &shorturlpb.GetUrlResponse{}, status.Errorf(codes.Internal, "unknown error: %s", err)
		}
		if serviceErr.Reason == service.ReasonNotFound {
			return &shorturlpb.GetUrlResponse{}, status.Errorf(codes.NotFound, "%s error: %s", serviceErr.Reason, serviceErr.Err)
		}
		return &shorturlpb.GetUrlResponse{}, status.Errorf(codes.Internal, "%s error: %s", serviceErr.Reason, serviceErr.Err)
	}
	return &shorturlpb.GetUrlResponse{OgUrl: res}, nil
}

// func (s *Server) Run() error {
// 	e := echo.New()
// 	apiGroup := e.Group("/api/v1")
// 	apiGroup.GET("/health-check", s.handleHealthCheck)
// 	return e.Start(s.listenAddr)
// }

// func (s *Server) handleHealthCheck(ctx echo.Context) error {
// 	return ctx.String(http.StatusOK, "OK")
// }
