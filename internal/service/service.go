package service

import "context"

type ShortUrlService interface {
	Shorten(ctx context.Context, someString string) (string, error)
	Reverse(ctx context.Context, shortLink string) (string, error)
}
