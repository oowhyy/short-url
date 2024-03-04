package postgres

import (
	"context"

	"github.com/uptrace/bun"
)

type PgStorage struct {
	db *bun.DB
}

func NewPgStorage(bunDB *bun.DB) *PgStorage {
	return &PgStorage{
		db: bunDB,
	}
}

func (ps *PgStorage) Save(ctx context.Context, shortKey string, longValue string) error {
	panic("not implemented") // TODO: Implement
}

func (ps *PgStorage) FindByKey(ctx context.Context, shortKey string) (string, bool, error) {
	panic("not implemented") // TODO: Implement
}

func (ps *PgStorage) FindByValue(ctx context.Context, longVlaue string) (string, bool, error) {
	panic("not implemented") // TODO: Implement
}
