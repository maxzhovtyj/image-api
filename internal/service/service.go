package service

import (
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/repository"
	"image"
)

type Images interface {
	Get(imageID uuid.UUID, quality int) (image.Image, error)
	Create(image image.Image) (uuid.UUID, error)
}

type Service struct {
	Images
}

func New(repository *repository.Repository) *Service {
	return &Service{
		Images: NewImagesService(repository.Images),
	}
}
