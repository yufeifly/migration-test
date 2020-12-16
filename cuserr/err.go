package cuserr

import "errors"

var (
	ErrBadParams    = errors.New("bad params")
	ErrBadAddress   = errors.New("bad address")
	ErrEmptyAddress = errors.New("empty address")
)
