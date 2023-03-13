package queue

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/service"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
	"image"
	"log"
)

type Consumer struct {
	client    *rabbitmq.Consumer
	service   service.Images
	qualities []int
}

func NewConsumer(consumer *rabbitmq.Consumer, images service.Images, imageQualities []int) *Consumer {
	return &Consumer{
		client:    consumer,
		service:   images,
		qualities: imageQualities,
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
				log.Fatal(err)
			}

			currWidth := decodedImage.Bounds().Dx()
			currHeight := decodedImage.Bounds().Dy()

			newUUID, err := uuid.NewUUID()
			if err != nil {
				log.Fatal(err)
			}

			for _, quality := range c.qualities {
				newWidth := currWidth * quality / 100
				newHeight := currHeight * quality / 100

				resizedImg, err := c.service.Resize(decodedImage, newWidth, newHeight)
				if err != nil {
					log.Fatal(err)
				}

				err = c.service.Create(resizedImg, newUUID, msg.ContentType, quality)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
