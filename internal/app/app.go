package app

import (
	"github.com/maxzhovtyj/image-api/config"
	delivery "github.com/maxzhovtyj/image-api/internal/delivery/http"
	"github.com/maxzhovtyj/image-api/internal/delivery/queue"
	"github.com/maxzhovtyj/image-api/internal/repository"
	"github.com/maxzhovtyj/image-api/internal/server"
	"github.com/maxzhovtyj/image-api/internal/service"
	"github.com/maxzhovtyj/image-api/pkg/img"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
	"log"
)

func Run(config *config.Config) {
	client, err := rabbitmq.NewClient(&config.AMQP)
	if err != nil {
		log.Fatal(err)
	}

	broker := rabbitmq.NewMessageBroker(client)
	defer broker.Publisher.CloseChan()

	imageManager := img.NewManager()

	repo := repository.New("img", imageManager)
	services := service.New(repo, broker.Publisher)
	handler := delivery.NewHandler(services)

	srv := server.New(config, handler.Init())

	imageQualities := []int{100, 75, 50, 25}

	queueConsumer := queue.NewConsumer(broker.Consumer, services.Images, imageQualities)

	err = queueConsumer.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
