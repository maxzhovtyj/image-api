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

	amqpUserEnv     = "AMQP_USER"
	amqpPasswordEnv = "AMQP_PASSWORD"
	amqpHostEnv     = "AMQP_HOST"
	amqpPortEnv     = "AMQP_PORT"
)

type (
	Config struct {
		HTTP
		AMQP
	}

	HTTP struct {
		Host string
		Port string
	}

	AMQP struct {
		Host     string
		Port     string
		User     string
		Password string
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

	cfg.AMQP.User = os.Getenv(amqpUserEnv)
	cfg.AMQP.Password = os.Getenv(amqpPasswordEnv)
	cfg.AMQP.Host = os.Getenv(amqpHostEnv)
	cfg.AMQP.Port = os.Getenv(amqpPortEnv)
}
