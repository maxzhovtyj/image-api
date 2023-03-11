package rabbitmq

import (
	"fmt"
	"github.com/maxzhovtyj/image-api/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewClient(cfg *config.AMQP) (*amqp.Connection, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)

	client, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}

	return client, err
}
