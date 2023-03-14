package queue

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/service"
	"github.com/maxzhovtyj/image-api/pkg/logger"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
	"image"
)

type Consumer struct {
	client    *rabbitmq.Consumer
	service   service.Images
	qualities []int
	logger    logger.Logger
}

func NewConsumer(
	consumer *rabbitmq.Consumer,
	images service.Images,
	imageQualities []int,
	logger logger.Logger,
) *Consumer {
	return &Consumer{
		client:    consumer,
		service:   images,
		qualities: imageQualities,
		logger:    logger,
	}
}

func (c *Consumer) Start() error {
	go c.worker()

	return nil
}

func (c *Consumer) worker() {
	messages := c.client.ConsumeMessages()

	for {
		select {
		case msg := <-messages:
			decodedImage, _, err := image.Decode(bytes.NewReader(msg.Body))
			if err != nil {
				c.logger.Errorf("failed to decode image due to %v", err)
				return
			}

			currWidth := decodedImage.Bounds().Dx()
			currHeight := decodedImage.Bounds().Dy()

			newUUID, err := uuid.NewUUID()
			if err != nil {
				c.logger.Errorf("failed create img uuid, %v", err)
				return
			}

			for _, quality := range c.qualities {
				newWidth := currWidth * quality / 100
				newHeight := currHeight * quality / 100

				resizedImg, err := c.service.Resize(decodedImage, newWidth, newHeight)
				if err != nil {
					c.logger.Errorf("failed to resize image due to %v", err)
					return
				}

				err = c.service.Create(resizedImg, newUUID, msg.ContentType, quality)
				if err != nil {
					c.logger.Errorf("failed to create and write new image due to %v", err)
					return
				}
			}
		}
	}
}
