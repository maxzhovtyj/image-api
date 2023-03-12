package repository

import (
	"github.com/google/uuid"
	"image"
)

type ImagesRepo struct {
	dir string
}

func NewImagesRepo(imageDirPath string) *ImagesRepo {
	return &ImagesRepo{
		dir: imageDirPath,
	}
}

func (r *ImagesRepo) Get(imageID uuid.UUID, quality int) (image.Image, error) {
	return nil, nil
}

func (r *ImagesRepo) Create(image image.Image) error {
	return nil
}
