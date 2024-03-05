package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"testing"

	"github.com/oowhyy/short-url/migrations"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/sync/errgroup"
)

const (
	testDNS = "postgresql://postgres:testpwd@localhost:5436/postgres?sslmode=disable"
)

var (
	once    sync.Once
	testBun *bun.DB
)

func newTestDb(t *testing.T) *PgStorage {
	once.Do(func() {
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(testDNS)))

		// Run all migrations before starting the app
		log.Info().Msg("migrating...")
		if err := migrations.Migrate(sqldb); err != nil {
			fmt.Printf("failed migrate: %v", err)
			return
		}
		bunDB := bun.NewDB(sqldb, pgdialect.New())
		if err := bunDB.Ping(); err != nil {
			fmt.Printf("failed to ping database: %v", err)
			t.FailNow()
			return
		}
		testBun = bunDB
	})
	if testBun == nil {
		t.FailNow()
		return nil
	}

	store := NewPgStorage(testBun)
	return store
}

func TestPgStorage_Save(t *testing.T) {
	t.Parallel()
	db := newTestDb(t)
	t.Cleanup(func() {
		if testBun == nil {
			return
		}
		for _, tt := range testDataSave {
			_, err := testBun.NewDelete().Table("shorturl").Where("short = ? ", tt.short).Exec(context.Background())
			if err != nil {
				log.Warn().Err(err).Msg("clean up error")
			}
		}
	})
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

func TestPgStorage_FindByKey(t *testing.T) {
	t.Parallel()
	db := newTestDb(t)
	t.Cleanup(func() {
		if testBun == nil {
			return
		}
		for _, tt := range testDataFindByKey {
			_, err := testBun.NewDelete().Table("shorturl").Where("short = ? ", tt.short).Exec(context.Background())
			if err != nil && err != sql.ErrNoRows {
				log.Warn().Err(err).Msg("clean up error")
			}
		}
	})
	// setup
	dataKeyValue := map[string]string{}
	dataNoSeed := map[string]bool{}
	for _, tt := range testDataFindByKey {
		if tt.seed {
			model := &ShortUrl{
				Short: tt.short,
				Long:  tt.val,
			}
			_, err := testBun.NewInsert().Model(model).Exec(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("seeding error")
				t.FailNow()
			}
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

func TestPgStorage_FindByValue(t *testing.T) {
	t.Parallel()
	db := newTestDb(t)
	t.Cleanup(func() {
		if testBun == nil {
			return
		}
		for _, tt := range testDataFindByValue {
			_, err := testBun.NewDelete().Table("shorturl").Where("short = ? ", tt.short).Exec(context.Background())
			if err != nil && err != sql.ErrNoRows {
				log.Warn().Err(err).Msg("clean up error")
			}
		}
	})
	// setup
	dataValueByKey := map[string]string{}
	dataNoSeed := map[string]bool{}
	for _, tt := range testDataFindByValue {
		if tt.seed {
			model := &ShortUrl{
				Short: tt.short,
				Long:  tt.val,
			}
			_, err := testBun.NewInsert().Model(model).Exec(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("seeding error")
				t.FailNow()
			}
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
