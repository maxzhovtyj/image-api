package queue

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/pkg/img"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
	"image"
	"log"
)

// TODO invalid messages consume

type Consumer struct {
	consumer     *rabbitmq.Consumer
	imageManager img.ImageManager
	qualities    []int
}

func NewConsumer(consumer *rabbitmq.Consumer, manager img.ImageManager) *Consumer {
	return &Consumer{
		consumer:     consumer,
		imageManager: manager,
		qualities:    []int{100, 75, 50, 25},
	}
}

func (c *Consumer) Start() error {
	go c.worker()

	return nil
}

func (c *Consumer) worker() {
	for {
		select {
		case msg := <-c.consumer.ConsumeMessages():
			decodedImage, _, err := image.Decode(bytes.NewReader(msg.Body))
			if err != nil {
				log.Fatal(err)
			}

			currWidth := decodedImage.Bounds().Dx()
			currHeight := decodedImage.Bounds().Dy()

			for _, quality := range c.qualities {
				newWidth := currWidth * quality / 100
				newHeight := currHeight * quality / 100

				resizedImage := c.imageManager.Resize(
					uint(newWidth),
					uint(newHeight),
					decodedImage,
				)

				imageUUID, err := uuid.NewUUID()
				if err != nil {
					log.Fatal(err)
				}

				err = c.imageManager.Write(
					fmt.Sprintf("img/%s_%d.jpg", imageUUID.String(), quality),
					resizedImage,
				)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
