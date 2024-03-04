package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oowhyy/short-url/internal/service"
)

type PostRequest struct {
	OgUrl string `json:"ogUrl"`
}

type PostResponse struct {
	ShortLink string `json:"shortLink"`
	Error     string `json:"error"`
}

type GetResponse struct {
	OgUrl string `json:"ogUrl"`
	Error string `json:"error"`
}

func (s *Server) handleShorten(ctx echo.Context) error {
	var reqBody PostRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, &PostResponse{
			ShortLink: "",
			Error:     "bad request params",
		})
	}
	res, err := s.shortUrlService.Shorten(ctx.Request().Context(),reqBody.OgUrl)
	if err != nil {
		var serviceErr *service.Error
		ok := errors.As(err, &serviceErr)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, &PostResponse{
				ShortLink: "",
				Error:     fmt.Sprintf("unknown: %s", err.Error()),
			})
		}
		status := http.StatusInternalServerError
		if serviceErr.Reason == service.ReasonInvalidReq {
			status = http.StatusBadRequest
		}
		return ctx.JSON(status, &PostResponse{
			ShortLink: "",
			Error:     serviceErr.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, &PostResponse{
		ShortLink: res,
		Error:     "",
	})
}

func (s *Server) handleReverse(ctx echo.Context) error {
	shortLink := ctx.Param("short")
	res, err := s.shortUrlService.Reverse(ctx.Request().Context(),shortLink)
	if err != nil {
		var serviceErr *service.Error
		ok := errors.As(err, &serviceErr)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, &GetResponse{
				OgUrl: "",
				Error: fmt.Sprintf("unknown: %s", err),
			})
		}
		status := http.StatusInternalServerError
		if serviceErr.Reason == service.ReasonNotFound {
			status = http.StatusNotFound
		}
		return ctx.JSON(status, &GetResponse{
			OgUrl: "",
			Error: serviceErr.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, &GetResponse{
		OgUrl: res,
		Error: "",
	})
}

func (s *Server) handleRedirect(ctx echo.Context) error {
	shortLink := ctx.Param("short")
	res, err := s.shortUrlService.Reverse(ctx.Request().Context(),shortLink)
	if err != nil {
		var serviceErr *service.Error
		ok := errors.As(err, &serviceErr)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, &GetResponse{
				OgUrl: "",
				Error: fmt.Sprintf("unknown: %s", err),
			})
		}
		status := http.StatusInternalServerError
		if serviceErr.Reason == service.ReasonNotFound {
			status = http.StatusNotFound
		}
		return ctx.JSON(status, &GetResponse{
			OgUrl: "",
			Error: serviceErr.Error(),
		})
	}
	return ctx.Redirect(http.StatusSeeOther, res)
}

func (s *Server) handleHealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}
