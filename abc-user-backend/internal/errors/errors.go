package apperrors

import "errors"

var (
	ErrEmailExists  = errors.New("email already exists")
	ErrInvalidAge   = errors.New("user must be at least 18")
	ErrUserNotFound = errors.New("user not found")
)
