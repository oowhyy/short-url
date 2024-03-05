package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oowhyy/short-url/internal/service"
	"github.com/oowhyy/short-url/pkg/shorturlpb"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	config          *Config
	logger          zerolog.Logger
	shortUrlService service.ShortUrlService

	shorturlpb.UnimplementedShortUrlServer
}

func NewServer(config *Config, logger zerolog.Logger, shortUrlService service.ShortUrlService) *Server {
	return &Server{
		config:          config,
		logger:          logger,
		shortUrlService: shortUrlService,
	}
}

func (s *Server) Run(ctx context.Context) error {
	wg, ctx := errgroup.WithContext(ctx)
	if s.config.ListenGrpc {
		wg.Go(func() error {
			err := s.RunGrpc(ctx)
			if err != nil {
				fmt.Println("grpc close err", err)
				return err
			}
			return nil
		})
	}
	if s.config.ListenHttp {
		wg.Go(func() error {
			err := s.RunHttp(ctx)
			if err != nil {
				fmt.Println("http close err", err)
				return err
			}
			return nil
		})
	}
	if err := wg.Wait(); err != nil {
		return err
	}
	return nil
}

func (s *Server) RunGrpc(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.config.ListenAddrGrpc)
	if err != nil {
		return fmt.Errorf("failed to listen tcp port: %w", err)
	}
	server := grpc.NewServer()
	shorturlpb.RegisterShortUrlServer(server, s)
	reflection.Register(server)
	s.logger.Info().Str("listenAddr", s.config.ListenAddrGrpc).Msg("running grpc server")
	go func() {
		<-ctx.Done()
		server.Stop()
	}()
	return server.Serve(lis)
}

func (s *Server) RunHttp(ctx context.Context) error {
	e := echo.New()
	e.HideBanner = true
	// api
	apiGroup := e.Group("/api/v1")
	apiGroup.GET("/health-check", s.handleHealthCheck)
	apiGroup.POST("/shorturl", s.handleShorten)
	apiGroup.GET("/shorturl/:short", s.handleReverse)
	// redirect
	e.GET("/:short", s.handleRedirect)

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		e.Shutdown(ctx)
	}()
	s.logger.Info().Str("listenAddr", s.config.ListenAddrHttp).Msg("running http server")
	err := e.Start(s.config.ListenAddrHttp)
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
