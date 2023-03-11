package domain

import (
	"mime/multipart"
)

type File struct {
	Name   string
	Size   int64
	Reader multipart.File
}
