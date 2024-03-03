package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/oowhyy/short-url/internal/config"
	"github.com/oowhyy/short-url/internal/server"
	"github.com/oowhyy/short-url/internal/service"
	"github.com/oowhyy/short-url/internal/storage/memory"
	"github.com/rs/zerolog"
)

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()
	cfg := config.MustLoadPath(*cfgPath)
	logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	baseLogger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(logLevel)
	store := memory.NewMemoryStorage()
	service := service.NewHasherService(cfg.Service, baseLogger.With().Str("component", "service").Logger(), store)

	server := server.NewServer(cfg.Server, baseLogger.With().Str("component", "server").Logger(), service)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	err = server.Run(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("MAIN EXIT")
}
