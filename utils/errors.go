package utils

import (
	"errors"
)

var (
	ErrInvalidUrl = errors.New("invalid url, must start with 'https://'")
	ErrEmptyBody  = errors.New("empty body")
)
