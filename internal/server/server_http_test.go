package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/oowhyy/short-url/internal/service"
	"github.com/oowhyy/short-url/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	ErrServiceInvalid = &service.Error{Reason: service.ReasonInvalidReq, Err: errors.New("some error")}
	ErrServiceInternal = &service.Error{Reason: service.ReasonService, Err: errors.New("some error")}
)

func TestServer_handleShorten(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		setupRequest  func(t *testing.T) *http.Request
		setupMock     func(t *testing.T) *mocks.ShortUrlService
		wantCode      int
		setupWantBody func(t *testing.T) string
	}{
		{
			name: "ok",
			setupRequest: func(t *testing.T) *http.Request {
				body, err := json.Marshal(&PostRequest{OgUrl: "ws://somevaliduri.com"})
				if err != nil {
					t.FailNow()
					return nil
				}
				return httptest.NewRequest(http.MethodPost, "/api/v1/service", bytes.NewReader(body))
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Shorten(mock.Anything).Return("mockoutput", nil).Once()
				return sus
			},
			wantCode: http.StatusOK,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(PostResponse{
					ShortLink: "mockoutput",
					Error:     "",
				})
				if err != nil {
					t.FailNow()
					return ""
				}
				return string(data)
			},
		},
		{
			name: "error bad json",
			setupRequest: func(t *testing.T) *http.Request {
				return httptest.NewRequest(http.MethodPost, "/api/v1/service", strings.NewReader("not a json"))
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				return sus
			},
			wantCode: http.StatusBadRequest,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(PostResponse{
					ShortLink: "",
					Error:     "bad request params",
				})
				if err != nil {
					t.FailNow()
					return ""
				}
				return string(data)
			},
		},
		{
			name: "error unknown",
			setupRequest: func(t *testing.T) *http.Request {
				body, err := json.Marshal(&PostRequest{OgUrl: "ws://somevaliduri.com"})
				if err != nil {
					t.FailNow()
					return nil
				}
				return httptest.NewRequest(http.MethodPost, "/api/v1/service", bytes.NewReader(body))
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Shorten(mock.Anything).Return("", errors.New("some untyped error")).Once()
				return sus
			},
			wantCode: http.StatusInternalServerError,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(PostResponse{
					ShortLink: "",
					Error:     "unknown: some untyped error",
				})
				if err != nil {
					t.FailNow()
					return ""
				}
				return string(data)
			},
		},
		{
			name: "error invalid uri",
			setupRequest: func(t *testing.T) *http.Request {
				body, err := json.Marshal(&PostRequest{OgUrl: "ws://somevaliduri.com"})
				if err != nil {
					t.FailNow()
					return nil
				}
				return httptest.NewRequest(http.MethodPost, "/api/v1/service", bytes.NewReader(body))
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Shorten(mock.Anything).Return("", ErrServiceInvalid).Once()
				return sus
			},
			wantCode: http.StatusBadRequest,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(PostResponse{
					ShortLink: "",
					Error:     ErrServiceInvalid.Error(),
				})
				if err != nil {
					t.FailNow()
					return ""
				}
				return string(data)
			},
		},
		{
			name: "error internal",
			setupRequest: func(t *testing.T) *http.Request {
				body, err := json.Marshal(&PostRequest{OgUrl: "ws://somevaliduri.com"})
				if err != nil {
					t.FailNow()
					return nil
				}
				return httptest.NewRequest(http.MethodPost, "/api/v1/service", bytes.NewReader(body))
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Shorten(mock.Anything).Return("", ErrServiceInternal).Once()
				return sus
			},
			wantCode: http.StatusInternalServerError,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(PostResponse{
					ShortLink: "",
					Error:     ErrServiceInternal.Error(),
				})
				if err != nil {
					t.FailNow()
					return ""
				}
				return string(data)
			},
			
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			server := newTestServer(tt.setupMock(t))
			e := echo.New()
			resp := httptest.NewRecorder()
			req := tt.setupRequest(t)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			ctx := e.NewContext(req, resp)
			err := server.handleShorten(ctx)
			require.NoError(t, err)
			require.Equal(t, tt.wantCode, resp.Code)
			wantBody := tt.setupWantBody(t)
			require.JSONEq(t, wantBody, resp.Body.String())
		})
	}
}
