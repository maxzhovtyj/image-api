package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/repository"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
	"image"
)

type Images interface {
	GetAll() ([]uuid.UUID, error)
	Get(imageID uuid.UUID, quality int) ([]byte, error)
	Create(image image.Image, imageID uuid.UUID, contentType string, quality int) error
	Resize(img image.Image, width, height int) (image.Image, error)
}

type Publisher interface {
	Publish(ctx context.Context, image []byte) error
}

type Service struct {
	Images
	Publisher
}

func New(repository *repository.Repository, publisher *rabbitmq.Publisher) *Service {
	return &Service{
		Images:    NewImagesService(repository.Images),
		Publisher: NewPublisherService(publisher),
	}
}
