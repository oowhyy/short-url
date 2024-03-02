package storage

type ShortUrlStorage interface {
	Save(shortKey, longValue string) error
	FindByKey(shortKey string) (string, bool, error)
	FindByValue(longVlaue string) (string, bool, error)
}
