package memory

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

var (
	testMemory *MemoryStorage
	once       sync.Once
)

func newTestMemory(t *testing.T) *MemoryStorage {
	once.Do(func() {
		mem := NewMemoryStorage()
		testMemory = mem
	})
	if testMemory == nil {
		t.FailNow()
		return nil
	}
	return testMemory
}

func TestMemoryStorage_Save(t *testing.T) {
	t.Parallel()
	db := newTestMemory(t)
	var wg errgroup.Group
	for _, tt := range testDataSave {
		tt := tt
		wg.Go(func() error {
			if err := db.Save(context.Background(), tt.short, tt.val); err != nil {
				return err
			}
			return nil
		})
	}
	err := wg.Wait()
	require.NoError(t, err)
}

func TestMemoryStorage_FindByKey(t *testing.T) {
	t.Parallel()
	db := newTestMemory(t)
	// setup
	dataKeyValue := map[string]string{}
	dataNoSeed := map[string]bool{}
	for _, tt := range testDataFindByKey {
		if tt.seed {
			db.keyValue.Store(tt.short, tt.val)
			db.valueKey.Store(tt.val, tt.short)
			dataKeyValue[tt.short] = tt.val
		} else {
			dataNoSeed[tt.val] = true
		}
	}
	// assert
	var wg errgroup.Group
	for _, tt := range testDataFindByKey {
		tt := tt
		wg.Go(func() error {
			long, ok, err := db.FindByKey(context.Background(), tt.short)
			require.NoError(t, err)
			if dataNoSeed[tt.val] {
				require.False(t, ok)
				require.Equal(t, "", long)
			} else {
				require.True(t, ok)
				require.Equal(t, dataKeyValue[tt.short], long)
			}
			return nil
		})
	}
	err := wg.Wait()
	require.NoError(t, err)
}

func TestMemoryStorage_FindByValue(t *testing.T) {
	t.Parallel()
	db := newTestMemory(t)
	// setup
	dataValueByKey := map[string]string{}
	dataNoSeed := map[string]bool{}
	for _, tt := range testDataFindByValue {
		if tt.seed {
			db.keyValue.Store(tt.short, tt.val)
			db.valueKey.Store(tt.val, tt.short)
			dataValueByKey[tt.val] = tt.short
		} else {
			dataNoSeed[tt.val] = true
		}
	}
	// assert
	var wg errgroup.Group
	for _, tt := range testDataFindByValue {
		tt := tt
		wg.Go(func() error {
			short, ok, err := db.FindByValue(context.Background(), tt.val)
			require.NoError(t, err)
			if dataNoSeed[tt.val] {
				require.False(t, ok)
				require.Equal(t, "", short)
			} else {
				require.True(t, ok)
				require.Equal(t, dataValueByKey[tt.val], short)
			}
			return nil
		})
	}
	err := wg.Wait()
	require.NoError(t, err)
}
