package memory

import (
	"context"
	"sync"
)

type MemoryStorage struct {
	keyValue sync.Map
	valueKey sync.Map
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (ms *MemoryStorage) Save(ctx context.Context, shortKey string, longValue string) error {
	ms.keyValue.Store(shortKey, longValue)
	ms.valueKey.Store(longValue, shortKey)
	return nil
}

func (ms *MemoryStorage) FindByKey(ctx context.Context, shortKey string) (string, bool, error) {
	val, ok := ms.keyValue.Load(shortKey)
	if !ok {
		return "", false, nil
	}
	return val.(string), true, nil
}

func (ms *MemoryStorage) FindByValue(ctx context.Context, longVlaue string) (string, bool, error) {
	key, ok := ms.valueKey.Load(longVlaue)
	if !ok {
		return "", false, nil
	}
	return key.(string), true, nil
}
