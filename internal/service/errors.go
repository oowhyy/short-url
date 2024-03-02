package service

import (
	"fmt"
)

const (
	ReasonInvalidReq = "invalid request"
	ReasonService    = "service"
	ReasonStorage    = "storage"
	ReasonNotFound   = "not found"
)

type Error struct {
	Reason string
	Err    error
}

func (e *Error) Unwrap() error { return e.Err }
func (e *Error) Error() string { return fmt.Sprintf("%s: %s", e.Reason, e.Err) }
