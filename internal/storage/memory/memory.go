package memory

import "sync"

type MemoryStorage struct {
	keyValue sync.Map
	valueKey sync.Map
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (ms *MemoryStorage) Save(shortKey string, longValue string) error {
	ms.keyValue.Store(shortKey, longValue)
	ms.valueKey.Store(longValue, shortKey)
	return nil
}

func (ms *MemoryStorage) FindByKey(shortKey string) (string, bool, error) {
	val, ok := ms.keyValue.Load(shortKey)
	if !ok {
		return "", false, nil
	}
	return val.(string), true, nil
}

func (ms *MemoryStorage) FindByValue(longVlaue string) (string, bool, error) {
	key, ok := ms.valueKey.Load(longVlaue)
	if !ok {
		return "", false, nil
	}
	return key.(string), true, nil
}
