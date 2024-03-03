package service

import (
	"errors"
	"testing"

	"github.com/oowhyy/short-url/internal/storage"
	"github.com/oowhyy/short-url/mocks"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	testBaseKey = "58e88fb0-eff2-4bd3-8901-b310910511c5"
)

func newTestHasher(store storage.ShortUrlStorage) *HasherService {
	cfg := &Config{
		BaseKey: testBaseKey,
	}
	return NewHasherService(cfg, zerolog.New(nil).Level(zerolog.Disabled), store)
}

func TestHasherService_Shorten(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		someString string
		setupMock  func(t *testing.T) *mocks.ShortUrlStorage
		want       string
		wantErr    *Error
	}{
		{
			name:       "ok hashed",
			someString: "https://example.com/abc?name=123#here",
			setupMock: func(t *testing.T) *mocks.ShortUrlStorage {
				store := mocks.NewShortUrlStorage(t)
				store.EXPECT().FindByValue(mock.Anything).Return("", false, nil).Once() // not found
				store.EXPECT().FindByKey(mock.Anything).Return("", false, nil).Once()   // can save
				store.EXPECT().Save(mock.Anything, mock.Anything).Return(nil).Once()    // save
				return store
			},
			want:    "nlGBTnTYkR",
			wantErr: nil,
		},
		{
			name:       "ok exists",
			someString: "ws://anyvaliduri.com:1111/abc?name=123#here",
			setupMock: func(t *testing.T) *mocks.ShortUrlStorage {
				store := mocks.NewShortUrlStorage(t)
				store.EXPECT().FindByValue(mock.Anything).Return("mockoutput", true, nil).Once() // exists
				return store
			},
			want:    "mockoutput",
			wantErr: nil,
		},
		{
			name:       "error invalid uri",
			someString: "google.com", // must be a request uri
			setupMock: func(t *testing.T) *mocks.ShortUrlStorage {
				store := mocks.NewShortUrlStorage(t)
				return store
			},
			want:    "",
			wantErr: &Error{ReasonInvalidReq, errors.New("some wrapped error")},
		},
		{
			name:       "error storage find by value",
			someString: "ws://anyvaliduri.com:1111/abc?name=123#here",
			setupMock: func(t *testing.T) *mocks.ShortUrlStorage {
				store := mocks.NewShortUrlStorage(t)
				store.EXPECT().FindByValue(mock.Anything).Return("", false, errors.New("storage find err"))
				return store
			},
			want:    "",
			wantErr: &Error{ReasonStorage, errors.New("some wrapped error")},
		},
		{
			name:       "error storage find by key",
			someString: "ws://anyvaliduri.com:1111/abc?name=123#here",
			setupMock: func(t *testing.T) *mocks.ShortUrlStorage {
				store := mocks.NewShortUrlStorage(t)
				store.EXPECT().FindByValue(mock.Anything).Return("", false, nil).Once() // not found
				store.EXPECT().FindByKey(mock.Anything).Return("", false, errors.New("storage find err"))
				return store
			},
			want:    "",
			wantErr: &Error{ReasonStorage, errors.New("some wrapped error")},
		},
		{
			name:       "error storage save",
			someString: "ws://anyvaliduri.com:1111/abc?name=123#here",
			setupMock: func(t *testing.T) *mocks.ShortUrlStorage {
				store := mocks.NewShortUrlStorage(t)
				store.EXPECT().FindByValue(mock.Anything).Return("", false, nil).Once()                           // not found
				store.EXPECT().FindByKey(mock.Anything).Return("", false, nil).Once()                             // can save
				store.EXPECT().Save(mock.Anything, mock.Anything).Return(errors.New("save storage error")).Once() // save error
				return store
			},
			want:    "",
			wantErr: &Error{ReasonStorage, errors.New("some wrapped error")},
		},
		{
			name:       "error many collisions",
			someString: "ws://anyvaliduri.com:1111/abc?name=123#here",
			setupMock: func(t *testing.T) *mocks.ShortUrlStorage {
				store := mocks.NewShortUrlStorage(t)
				store.EXPECT().FindByValue(mock.Anything).Return("", false, nil).Once()        // not found
				store.EXPECT().FindByKey(mock.Anything).Return("http://exists.com", true, nil) // always collide
				return store
			},
			want:    "",
			wantErr: &Error{ReasonService, errors.New("some wrapped error")},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			store := tt.setupMock(t)

			hasher := newTestHasher(store)
			got, err := hasher.Shorten(tt.someString)
			require.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				var customErr *Error
				ok := errors.As(err, &customErr)
				require.True(t, ok)
				require.Equal(t, tt.wantErr.Reason, customErr.Reason)
			}
		})
	}
}

func TestHasherService_Reverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		shortLink string
		setupMock func(t *testing.T) *mocks.ShortUrlStorage
		want      string
		wantErr   *Error
	}{
		{
			name:      "ok",
			shortLink: "mockinput",
			setupMock: func(t *testing.T) *mocks.ShortUrlStorage {
				store := mocks.NewShortUrlStorage(t)
				store.EXPECT().FindByKey("mockinput").Return("http://ogurl.com", true, nil).Once()
				return store
			},
			want:    "http://ogurl.com",
			wantErr: nil,
		},
		{
			name:      "error not found",
			shortLink: "mockinput",
			setupMock: func(t *testing.T) *mocks.ShortUrlStorage {
				store := mocks.NewShortUrlStorage(t)
				store.EXPECT().FindByKey("mockinput").Return("", false, nil).Once()
				return store
			},
			want:    "",
			wantErr: &Error{ReasonNotFound, errors.New("some wrapped error")},
		},
		{
			name:      "error not found",
			shortLink: "mockinput",
			setupMock: func(t *testing.T) *mocks.ShortUrlStorage {
				store := mocks.NewShortUrlStorage(t)
				store.EXPECT().FindByKey("mockinput").Return("", false, errors.New("storage save error")).Once()
				return store
			},
			want:    "",
			wantErr: &Error{ReasonStorage, errors.New("some wrapped error")},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			store := tt.setupMock(t)
			hasher := newTestHasher(store)
			got, err := hasher.Reverse(tt.shortLink)
			require.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				var customErr *Error
				ok := errors.As(err, &customErr)
				require.True(t, ok)
				require.Equal(t, tt.wantErr.Reason, customErr.Reason)
			}
		})
	}
}
