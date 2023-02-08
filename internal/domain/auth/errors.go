package auth

import "errors"

var (
	ErrPassIncorrect = errors.New("password mismatch")
	ErrEmptySecret   = errors.New("empty secret key for token creation")
	ErrZeroDuration  = errors.New("zero duration for token creation")
)
