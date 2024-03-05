package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/oowhyy/short-url/internal/config"
	"github.com/oowhyy/short-url/internal/conn"
	"github.com/oowhyy/short-url/internal/server"
	"github.com/oowhyy/short-url/internal/service"
	"github.com/oowhyy/short-url/internal/storage"
	"github.com/oowhyy/short-url/internal/storage/memory"
	"github.com/oowhyy/short-url/internal/storage/postgres"
	"github.com/rs/zerolog"
)

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()
	cfg := config.MustLoadPath(*cfgPath)
	logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	store := mustStorageFromType(cfg.StorageType)
	baseLogger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(logLevel)
	service := service.NewHasherService(cfg.Service, baseLogger.With().Str("component", "service").Logger(), store)

	server := server.NewServer(cfg.Server, baseLogger.With().Str("component", "server").Logger(), service)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	err = server.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MAIN EXIT")
}

func mustStorageFromType(sType string) storage.ShortUrlStorage {
	switch sType {
	case "memory":
		return memory.NewMemoryStorage()
	case "postgres":
		bunDb, err := conn.NewBunPostgres()
		if err != nil {
			log.Fatal(err.Error())
		}
		pg := postgres.NewPgStorage(bunDb)
		return pg
	default:
		log.Fatal("unknown storage type:", sType)
	}
	return nil
}
