package storage

import "context"

type ShortUrlStorage interface {
	Save(ctx context.Context, shortKey, longValue string) error
	FindByKey(ctx context.Context, shortKey string) (string, bool, error)
	FindByValue(ctx context.Context, longVlaue string) (string, bool, error)
}
