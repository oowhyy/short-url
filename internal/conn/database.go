package conn

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/oowhyy/short-url/migrations"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewBunPostgres() (*bun.DB, error) {
	pgDatabase := os.Getenv("POSTGRES_DB")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgAddr := fmt.Sprintf("%s:%s", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_CONTAINER_PORT"))
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithAddr(pgAddr),
		pgdriver.WithUser(pgUser),
		pgdriver.WithPassword(pgPassword),
		pgdriver.WithDatabase(pgDatabase),
		pgdriver.WithInsecure(true),
	))
	db := bun.NewDB(sqldb, pgdialect.New())
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}
	// Run all migrations before starting the app
	log.Info().Msg("migrating...")
	if err := migrations.Migrate(sqldb); err != nil {
		return nil, err
	}
	return db, nil
}
