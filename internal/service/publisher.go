package service

import (
	"bytes"
	"context"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
	"image"
	"image/jpeg"
)

type PublisherService struct {
	publisher *rabbitmq.Publisher
}

func NewPublisherService(publisher *rabbitmq.Publisher) *PublisherService {
	return &PublisherService{publisher: publisher}
}

func (s *PublisherService) Publish(ctx context.Context, image image.Image) error {
	buff := new(bytes.Buffer)

	err := jpeg.Encode(buff, image, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}

	err = s.publisher.PublishMessage(
		ctx,
		buff.Bytes(),
		"images/jpeg",
	)
	if err != nil {
		return err
	}

	return nil
}
