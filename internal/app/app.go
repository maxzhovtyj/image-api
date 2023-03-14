package app

import (
	"github.com/maxzhovtyj/image-api/config"
	delivery "github.com/maxzhovtyj/image-api/internal/delivery/http"
	"github.com/maxzhovtyj/image-api/internal/delivery/queue"
	"github.com/maxzhovtyj/image-api/internal/repository"
	"github.com/maxzhovtyj/image-api/internal/server"
	"github.com/maxzhovtyj/image-api/internal/service"
	"github.com/maxzhovtyj/image-api/pkg/img"
	"github.com/maxzhovtyj/image-api/pkg/logger"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
)

func Run(config *config.Config, logger logger.Logger) {
	logger.Info("initializing rabbitmq client")
	client, err := rabbitmq.NewClient(&config.AMQP)
	if err != nil {
		logger.Fatalf("failed to init rabbitmq client due to %v", err)
	}

	logger.Info("initializing new message broker")
	broker := rabbitmq.NewMessageBroker(client)
	defer broker.Publisher.CloseChan()

	imageManager := img.NewManager()

	logger.Info("initializing repository, services and handlers")
	repo := repository.New("img", imageManager)
	services := service.New(repo, broker.Publisher)
	handler := delivery.NewHandler(services, logger)

	logger.Info("initializing new server")
	srv := server.New(config, handler.Init())

	imageQualities := []int{100, 75, 50, 25}

	logger.Info("init and start new message consumer")
	queueConsumer := queue.NewConsumer(broker.Consumer, services.Images, imageQualities, logger)

	err = queueConsumer.Start()
	if err != nil {
		logger.Fatal("failed to start new queue consumer due to %v", err)
	}

	logger.Infof("Start application on port %s", config.HTTP.Port)
	err = srv.Run()
	if err != nil {
		logger.Fatal("failed to run application due to %v", err)
	}
}
