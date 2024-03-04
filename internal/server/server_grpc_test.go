package server

import (
	"context"
	"errors"
	"testing"

	"github.com/oowhyy/short-url/internal/service"
	"github.com/oowhyy/short-url/mocks"
	"github.com/oowhyy/short-url/pkg/shorturlpb"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newTestServer(svc service.ShortUrlService) *Server {
	cfg := &Config{}
	return NewServer(cfg, zerolog.New(nil).Level(zerolog.Disabled), svc)
}

func TestServer_Shorten(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		req       *shorturlpb.PostUrlRequest
		setupMock func(t *testing.T) *mocks.ShortUrlService
		want      *shorturlpb.PostUrlResponse
		wantErr   error
	}{
		{
			name: "ok",
			req: &shorturlpb.PostUrlRequest{
				OgUrl: "anything",
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Shorten(mock.Anything,mock.Anything).Return("mockoutput", nil).Once()
				return sus
			},
			want: &shorturlpb.PostUrlResponse{
				ShortLink: "mockoutput",
			},
			wantErr: nil,
		},
		{
			name: "error unknown err",
			req: &shorturlpb.PostUrlRequest{
				OgUrl: "anything",
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Shorten(mock.Anything,mock.Anything).Return("", errors.New("some untyped error")).Once()
				return sus
			},
			want:    &shorturlpb.PostUrlResponse{},
			wantErr: status.Error(codes.Internal, "any err msg"),
		},
		{
			name: "error invalid",
			req: &shorturlpb.PostUrlRequest{
				OgUrl: "anything",
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Shorten(mock.Anything,mock.Anything).Return("", &service.Error{Reason: service.ReasonInvalidReq, Err: errors.New("some error")}).Once()
				return sus
			},
			want:    &shorturlpb.PostUrlResponse{},
			wantErr: status.Error(codes.InvalidArgument, "any err msg"),
		},
		{
			name: "error internal",
			req: &shorturlpb.PostUrlRequest{
				OgUrl: "anything",
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Shorten(mock.Anything,mock.Anything).Return("", &service.Error{Reason: service.ReasonService, Err: errors.New("some error")}).Once()
				return sus
			},
			want:    &shorturlpb.PostUrlResponse{},
			wantErr: status.Error(codes.Internal, "any err msg"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			server := newTestServer(tt.setupMock(t))
			got, err := server.Shorten(context.Background(), tt.req)
			require.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				wantCode := status.Code(tt.wantErr)
				gotCode := status.Code(err)
				require.Equal(t, wantCode, gotCode)
			}
		})
	}
}

func TestServer_Reverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		req       *shorturlpb.GetUrlRequest
		setupMock func(t *testing.T) *mocks.ShortUrlService
		want      *shorturlpb.GetUrlResponse
		wantErr   error
	}{
		{
			name: "ok",
			req: &shorturlpb.GetUrlRequest{
				ShortLink: "anything",
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,mock.Anything).Return("mockoutput", nil).Once()
				return sus
			},
			want: &shorturlpb.GetUrlResponse{
				OgUrl: "mockoutput",
			},
			wantErr: nil,
		},
		{
			name: "error unknown",
			req: &shorturlpb.GetUrlRequest{
				ShortLink: "anything",
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,mock.Anything).Return("", errors.New("untyped error")).Once()
				return sus
			},
			want: &shorturlpb.GetUrlResponse{},
			wantErr: status.Error(codes.Internal, "some msg"),
		},
		{
			name: "error not found",
			req: &shorturlpb.GetUrlRequest{
				ShortLink: "anything",
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,mock.Anything).Return("", &service.Error{Reason: service.ReasonNotFound, Err: errors.New("some err")}).Once()
				return sus
			},
			want: &shorturlpb.GetUrlResponse{},
			wantErr: status.Error(codes.NotFound, "some msg"),
		},
		{
			name: "error internal",
			req: &shorturlpb.GetUrlRequest{
				ShortLink: "anything",
			},
			setupMock: func(t *testing.T) *mocks.ShortUrlService {
				sus := mocks.NewShortUrlService(t)
				sus.EXPECT().Reverse(mock.Anything,mock.Anything).Return("", &service.Error{Reason: service.ReasonService, Err: errors.New("some err")}).Once()
				return sus
			},
			want: &shorturlpb.GetUrlResponse{},
			wantErr: status.Error(codes.Internal, "some msg"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			server := newTestServer(tt.setupMock(t))
			got, err := server.Reverse(context.Background(), tt.req)
			require.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				wantCode := status.Code(tt.wantErr)
				gotCode := status.Code(err)
				require.Equal(t, wantCode, gotCode)
			}
		})
	}
}
