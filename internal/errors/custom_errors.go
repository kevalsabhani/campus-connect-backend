package errors

import "errors"

var (
	ErrEmptyConfig       = errors.New("one or more empty config")
	ErrInvalidPort       = errors.New("invalid port")
	ErrInvalidExpiration = errors.New("invalid expiration")
)
