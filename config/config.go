package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

const (
	httpHostEnv = "HTTP_HOST"
	httpPortEnv = "HTTP_PORT"
)

type (
	Config struct {
		HTTP
	}

	HTTP struct {
		Host string
		Port string
	}
)

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}

		setFromEnv(&config)
	})

	return &config
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv(httpHostEnv)
	cfg.HTTP.Port = os.Getenv(httpPortEnv)
}
