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
	TestErrServiceInvalid  = &service.Error{Reason: service.ReasonInvalidReq, Err: errors.New("some error")}
	TestErrServiceInternal = &service.Error{Reason: service.ReasonService, Err: errors.New("some error")}
	TestErrNotFound        = &service.Error{Reason: service.ReasonNotFound, Err: errors.New("some error")}
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
				sus.EXPECT().Shorten(mock.Anything,mock.Anything).Return("mockoutput", nil).Once()
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
					Error:     `invalid json body or header "Content-Type" is not set to "application/json"`,
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
				sus.EXPECT().Shorten(mock.Anything,mock.Anything).Return("", errors.New("some untyped error")).Once()
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
				sus.EXPECT().Shorten(mock.Anything,mock.Anything).Return("", TestErrServiceInvalid).Once()
				return sus
			},
			wantCode: http.StatusBadRequest,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(PostResponse{
					ShortLink: "",
					Error:     TestErrServiceInvalid.Error(),
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
				sus.EXPECT().Shorten(mock.Anything,mock.Anything).Return("", TestErrServiceInternal).Once()
				return sus
			},
			wantCode: http.StatusInternalServerError,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(PostResponse{
					ShortLink: "",
					Error:     TestErrServiceInternal.Error(),
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

func TestServer_handleReverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		setupContext  func(t *testing.T, resp http.ResponseWriter) echo.Context
		setupMock     func(t *testing.T) *mocks.ShortUrlService
		wantCode      int
		setupWantBody func(t *testing.T) string
	}{
		{
			name: "ok",
			setupContext: func(t *testing.T, resp http.ResponseWriter) echo.Context {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				ctx := e.NewContext(req, resp)
				ctx.SetPath("/api/v1/service/:short")
				ctx.SetParamNames("short")
				ctx.SetParamValues("_SHORTLINK")
				return ctx
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,"_SHORTLINK").Return("mockoutputfullurl", nil).Once()
				return sus
			},
			wantCode: http.StatusOK,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(GetResponse{OgUrl: "mockoutputfullurl", Error: ""})
				if err != nil {
					t.FailNow()
					return ""
				}
				return string(data)
			},
		},
		{
			name: "error unknown",
			setupContext: func(t *testing.T, resp http.ResponseWriter) echo.Context {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				ctx := e.NewContext(req, resp)
				ctx.SetPath("/api/v1/service/:short")
				ctx.SetParamNames("short")
				ctx.SetParamValues("_SHORTLINK")
				return ctx
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,"_SHORTLINK").Return("", errors.New("some error")).Once()
				return sus
			},
			wantCode: http.StatusInternalServerError,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(GetResponse{OgUrl: "", Error: "unknown: some error"})
				if err != nil {
					t.FailNow()
					return ""
				}
				return string(data)
			},
		},
		{
			name: "error not found",
			setupContext: func(t *testing.T, resp http.ResponseWriter) echo.Context {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				ctx := e.NewContext(req, resp)
				ctx.SetPath("/api/v1/service/:short")
				ctx.SetParamNames("short")
				ctx.SetParamValues("_SHORTLINK")
				return ctx
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,"_SHORTLINK").Return("", TestErrNotFound).Once()
				return sus
			},
			wantCode: http.StatusNotFound,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(GetResponse{OgUrl: "", Error: TestErrNotFound.Error()})
				if err != nil {
					t.FailNow()
					return ""
				}
				return string(data)
			},
		},
		{
			name: "error internal",
			setupContext: func(t *testing.T, resp http.ResponseWriter) echo.Context {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				ctx := e.NewContext(req, resp)
				ctx.SetPath("/api/v1/service/:short")
				ctx.SetParamNames("short")
				ctx.SetParamValues("_SHORTLINK")
				return ctx
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,"_SHORTLINK").Return("", TestErrServiceInternal).Once()
				return sus
			},
			wantCode: http.StatusInternalServerError,
			setupWantBody: func(t *testing.T) string {
				data, err := json.Marshal(GetResponse{OgUrl: "", Error: TestErrServiceInternal.Error()})
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
			resp := httptest.NewRecorder()
			ctx := tt.setupContext(t, resp)
			err := server.handleReverse(ctx)
			require.NoError(t, err)
			require.Equal(t, tt.wantCode, resp.Code)
			wantBody := tt.setupWantBody(t)
			require.JSONEq(t, wantBody, resp.Body.String())
		})
	}
}

func TestServer_handleRedirect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		setupContext func(t *testing.T, resp http.ResponseWriter) echo.Context
		setupMock    func(t *testing.T) *mocks.ShortUrlService
		wantCode     int
	}{
		{
			name: "ok",
			setupContext: func(t *testing.T, resp http.ResponseWriter) echo.Context {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				ctx := e.NewContext(req, resp)
				ctx.SetPath("/:short")
				ctx.SetParamNames("short")
				ctx.SetParamValues("_SHORTLINK")
				return ctx
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,"_SHORTLINK").Return("https://example.com", nil).Once()
				return sus
			},
			wantCode: http.StatusSeeOther,
		},
		{
			name: "error unknown",
			setupContext: func(t *testing.T, resp http.ResponseWriter) echo.Context {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				ctx := e.NewContext(req, resp)
				ctx.SetPath("/:short")
				ctx.SetParamNames("short")
				ctx.SetParamValues("_SHORTLINK")
				return ctx
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,"_SHORTLINK").Return("", errors.New("some error")).Once()
				return sus
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "error not found",
			setupContext: func(t *testing.T, resp http.ResponseWriter) echo.Context {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				ctx := e.NewContext(req, resp)
				ctx.SetPath("/:short")
				ctx.SetParamNames("short")
				ctx.SetParamValues("_SHORTLINK")
				return ctx
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,"_SHORTLINK").Return("", TestErrNotFound).Once()
				return sus
			},
			wantCode: http.StatusNotFound,
		},
		{
			name: "error internal",
			setupContext: func(t *testing.T, resp http.ResponseWriter) echo.Context {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				ctx := e.NewContext(req, resp)
				ctx.SetPath("/:short")
				ctx.SetParamNames("short")
				ctx.SetParamValues("_SHORTLINK")
				return ctx
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,"_SHORTLINK").Return("", TestErrServiceInternal).Once()
				return sus
			},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			server := newTestServer(tt.setupMock(t))
			resp := httptest.NewRecorder()
			ctx := tt.setupContext(t, resp)
			err := server.handleRedirect(ctx)
			require.NoError(t, err)
			require.Equal(t, tt.wantCode, resp.Code)
		})
	}
}

func TestServer_handleHealthCheck(t *testing.T) {
	t.Parallel()
	server := newTestServer(mocks.NewShortUrlService(t))
	resp := httptest.NewRecorder()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/health-check", nil)
	ctx := e.NewContext(req, resp)
	err := server.handleHealthCheck(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.Code)
}
