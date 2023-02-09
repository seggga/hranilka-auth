package auth

import "errors"

var (
	ErrPassIncorrect = errors.New("password mismatch")
	ErrEmptySecret   = errors.New("empty secret key for token creation")
	ErrZeroDuration  = errors.New("zero duration for token creation")
	ErrEmptyLogin    = errors.New("cannot create a user with empty login")
	ErrPassTooLong   = errors.New("given password is too long")
)
