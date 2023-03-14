package utils

import (
	"errors"
)

var (
	ErrUrlExists  = errors.New("url already exists in database")
	ErrInvalidUrl = errors.New("invalid url, must start with 'https://'")
	ErrEmptyBody  = errors.New("empty body")
	ErrInvalidID  = errors.New("incorrect short url")
	ErrEmptyID    = errors.New("no created short url")
)
