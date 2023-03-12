package service

import (
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/repository"
	"image"
)

type ImagesService struct {
	repo repository.Images
}

func NewImagesService(images repository.Images) *ImagesService {
	return &ImagesService{repo: images}
}

func (s *ImagesService) Get(imageID uuid.UUID, quality int) (image.Image, error) {
	return nil, nil
}

func (s *ImagesService) Create(image image.Image) (uuid.UUID, error) {
	return [16]byte{}, nil
}
