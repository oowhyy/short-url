package memory

type MemoryStorage struct {
}

func (ms *MemoryStorage) Save(shortKey string, longValue string) error {
	panic("not implemented") // TODO: Implement
}

func (ms *MemoryStorage) FindByKey(shortKey string) (string, bool, error) {
	panic("not implemented") // TODO: Implement
}

func (ms *MemoryStorage) FindByValue(longVlaue string) (string, bool, error) {
	panic("not implemented") // TODO: Implement
}
