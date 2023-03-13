package queue

import (
	"bytes"
	"github.com/maxzhovtyj/image-api/internal/service"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
	"image"
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
				return
			}

			currWidth := decodedImage.Bounds().Dx()
			currHeight := decodedImage.Bounds().Dy()

			for _, quality := range c.qualities {
				newWidth := currWidth * quality / 100
				newHeight := currHeight * quality / 100

				resizedImg := c.service.Resize(decodedImage, newWidth, newHeight)
				if err != nil {
					return
				}

				err = c.service.Create(resizedImg, "jpg", quality)
				if err != nil {
					return
				}
			}
		}
	}
}
