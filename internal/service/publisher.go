package service

import (
	"context"
	"github.com/maxzhovtyj/image-api/internal/domain"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
	"net/http"
)

type PublisherService struct {
	publisher *rabbitmq.Publisher
}

func NewPublisherService(publisher *rabbitmq.Publisher) *PublisherService {
	return &PublisherService{publisher: publisher}
}

func (s *PublisherService) Publish(ctx context.Context, image []byte) error {
	contentType := http.DetectContentType(image)

	switch contentType {
	case "image/jpeg", "image/jpg", "image/png":
		err := s.publisher.PublishMessage(
			ctx,
			image,
			contentType,
		)
		if err != nil {
			return err
		}
	default:
		return domain.ErrInvalidContentType
	}

	return nil
}
