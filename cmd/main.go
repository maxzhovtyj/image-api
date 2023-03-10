package main

import (
	"github.com/maxzhovtyj/image-api/config"
	"github.com/maxzhovtyj/image-api/internal/app"
)

func main() {
	cfg := config.Get()

	app.Run(cfg)
}
