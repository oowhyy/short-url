package service

import (
	"errors"
	"fmt"
	"hash/fnv"
	"net/url"

	"github.com/google/uuid"
	"github.com/oowhyy/short-url/internal/storage"
	"github.com/rs/zerolog"
)

const (
	alphabet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz" // base 53
	alphabetBase = uint64(len(alphabet))
	hashLength   = 10
)

type HasherService struct {
	logger          zerolog.Logger
	HashKey         string
	shortUrlStorage storage.ShortUrlStorage
}

func NewHasherService(config *Config, logger zerolog.Logger, shortUrlStorage storage.ShortUrlStorage) *HasherService {
	return &HasherService{
		logger:          logger,
		HashKey:         config.BaseKey,
		shortUrlStorage: shortUrlStorage,
	}
}

func (hs *HasherService) Shorten(someString string) (string, error) {
	parsedUrl, err := url.ParseRequestURI(someString)
	if err != nil {
		return "", &Error{ReasonInvalidReq, fmt.Errorf("parse url: %w", err)}
	}
	urlString := parsedUrl.String()
	long, ok, err := hs.shortUrlStorage.FindByValue(urlString)
	if err != nil {
		return "", &Error{ReasonStorage, fmt.Errorf("find by value: %w", err)}
	}
	if ok {
		return long, nil
	}
	for try := 0; try < 100; try++ {
		short := magicHash(hs.HashKey, urlString)
		// check collision
		_, ok, err := hs.shortUrlStorage.FindByKey(short)
		if err != nil {
			return "", &Error{ReasonStorage, fmt.Errorf("find by key: %w", err)}
		}
		if ok {
			// generate new key and retry
			hs.HashKey = uuid.NewString()
			hs.logger.Warn().Str("newKey", hs.HashKey).Msg("rare collision event - generating new key")
			continue
		}
		err = hs.shortUrlStorage.Save(short, urlString)
		if err != nil {
			return "", &Error{ReasonStorage, fmt.Errorf("save: %w", err)}
		}
		return short, nil
	}
	return "", &Error{ReasonService, errors.New("too many collisions")}
}

func (hs *HasherService) Reverse(shortLink string) (string, error) {
	long, ok, err := hs.shortUrlStorage.FindByKey(shortLink)
	if err != nil {
		return "", &Error{ReasonStorage, fmt.Errorf("find by key: %w", err)}
	}
	if !ok {
		return "", &Error{ReasonNotFound, fmt.Errorf("%q not found", shortLink)}
	}
	return long, nil
}

func magicHash(key string, someString string) string {
	toHash := []byte(key + someString)
	h := fnv.New64a()
	h.Write(toHash)
	sum := h.Sum64()
	result := make([]byte, hashLength)
	for i := range result {
		result[i] = alphabet[sum%alphabetBase]
		sum /= alphabetBase
	}
	return string(result)
}
