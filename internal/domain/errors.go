package domain

import "errors"

var (
	ErrInvalidContentType = errors.New("invalid content type")
	ErrImageNotFound      = errors.New("image was not found")
)
