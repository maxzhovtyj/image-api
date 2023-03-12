package repository

import (
	"github.com/google/uuid"
	"image"
)

type Images interface {
	Get(imageID uuid.UUID, quality int) (image.Image, error)
	Create(image image.Image) error
}

type Repository struct {
	Images
}

func New(imgDirPath string) *Repository {
	return &Repository{
		Images: NewImagesRepo(imgDirPath),
	}
}
