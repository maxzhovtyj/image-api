package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/repository"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
	"image"
)

type Images interface {
	Get(imageID uuid.UUID, quality int) (image.Image, error)
	Create(image image.Image, extension string, quality int) error
	Resize(img image.Image, width, height int) image.Image
}

type Publisher interface {
	Publish(ctx context.Context, image image.Image) error
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
