package server

import (
	"context"
	"errors"

	"github.com/oowhyy/short-url/internal/service"
	"github.com/oowhyy/short-url/pkg/shorturlpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Shorten(ctx context.Context, req *shorturlpb.PostUrlRequest) (*shorturlpb.PostUrlResponse, error) {
	someString := req.GetOgUrl()
	res, err := s.shortUrlService.Shorten(ctx, someString)
	if err != nil {
		var serviceErr *service.Error
		ok := errors.As(err, &serviceErr)
		if !ok {
			return &shorturlpb.PostUrlResponse{}, status.Errorf(codes.Internal, "unknown: %s", err)
		}
		if serviceErr.Reason == service.ReasonInvalidReq {
			return &shorturlpb.PostUrlResponse{}, status.Error(codes.InvalidArgument, serviceErr.Error())
		}
		return &shorturlpb.PostUrlResponse{}, status.Error(codes.Internal, serviceErr.Error())
	}
	return &shorturlpb.PostUrlResponse{ShortLink: res}, nil
}

func (s *Server) Reverse(ctx context.Context, req *shorturlpb.GetUrlRequest) (*shorturlpb.GetUrlResponse, error) {
	short := req.GetShortLink()
	res, err := s.shortUrlService.Reverse(ctx,short)
	if err != nil {
		var serviceErr *service.Error
		ok := errors.As(err, &serviceErr)
		if !ok {
			return &shorturlpb.GetUrlResponse{}, status.Errorf(codes.Internal, "unknown: %s", err)
		}
		if serviceErr.Reason == service.ReasonNotFound {
			return &shorturlpb.GetUrlResponse{}, status.Error(codes.NotFound, serviceErr.Error())
		}
		return &shorturlpb.GetUrlResponse{}, status.Error(codes.Internal, serviceErr.Error())
	}
	return &shorturlpb.GetUrlResponse{OgUrl: res}, nil
}
