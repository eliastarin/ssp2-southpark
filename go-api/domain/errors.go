package domain

import "errors"

var (
	ErrInvalidAuthor = errors.New("author is required")
	ErrInvalidBody   = errors.New("body is required")
)
