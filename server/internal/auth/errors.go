package auth

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidCredential = errors.New("invalid email or password")
	ErrEmailAlreadyExist = errors.New("email already exists")
)
