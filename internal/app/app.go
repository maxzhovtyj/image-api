package app

import (
	"github.com/maxzhovtyj/image-api/config"
	delivery "github.com/maxzhovtyj/image-api/internal/delivery/http"
	"github.com/maxzhovtyj/image-api/internal/repository"
	"github.com/maxzhovtyj/image-api/internal/server"
	"github.com/maxzhovtyj/image-api/internal/service"
	"github.com/maxzhovtyj/image-api/pkg/queue/rabbitmq"
	"log"
)

func Run(config *config.Config) {
	_, err := rabbitmq.NewClient(&config.AMQP)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New("/img")
	services := service.New(repo)
	handler := delivery.NewHandler(services)

	srv := server.New(config, handler.Init())

	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
