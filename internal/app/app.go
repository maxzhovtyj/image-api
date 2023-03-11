package app

import (
	"github.com/maxzhovtyj/image-api/config"
	delivery "github.com/maxzhovtyj/image-api/internal/delivery/http"
	"github.com/maxzhovtyj/image-api/internal/server"
	"github.com/maxzhovtyj/image-api/pkg/rabbitmq"
	"log"
)

func Run(config *config.Config) {
	handler := delivery.NewHandler()

	_, err := rabbitmq.NewClient(&config.AMQP)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(config, handler.Init())

	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
