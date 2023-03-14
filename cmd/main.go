package main

import (
	"github.com/maxzhovtyj/image-api/config"
	"github.com/maxzhovtyj/image-api/internal/app"
	"github.com/maxzhovtyj/image-api/pkg/logger"
	"log"
)

func main() {
	cfg := config.Get()

	appLogger := logger.NewAppLogger(cfg)
	err := appLogger.InitLogger()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg, appLogger)
}
