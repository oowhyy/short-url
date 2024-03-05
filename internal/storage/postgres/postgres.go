package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"
)

type ShortUrl struct {
	bun.BaseModel `bun:"table:shorturl"`

	Short string `bun:"short"`
	Long  string `bun:"long"`
}

type PgStorage struct {
	db *bun.DB
}

func NewPgStorage(bunDB *bun.DB) *PgStorage {
	return &PgStorage{
		db: bunDB,
	}
}

func (ps *PgStorage) Save(ctx context.Context, shortKey string, longValue string) error {
	model := &ShortUrl{
		Short: shortKey,
		Long:  longValue,
	}
	_, err := ps.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PgStorage) FindByKey(ctx context.Context, shortKey string) (string, bool, error) {
	var model ShortUrl
	err := ps.db.NewSelect().Model(&model).Where("short = ?", shortKey).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return model.Long, true, nil
}

func (ps *PgStorage) FindByValue(ctx context.Context, longVlaue string) (string, bool, error) {
	var model ShortUrl
	err := ps.db.NewSelect().Model(&model).Where("long = ?", longVlaue).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return model.Short, true, nil
}
