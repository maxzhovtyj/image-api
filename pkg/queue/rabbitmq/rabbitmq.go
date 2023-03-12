package rabbitmq

import (
	"context"
	"fmt"
	"github.com/maxzhovtyj/image-api/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	Body        []byte
	ContentType string
}

func NewClient(cfg *config.AMQP) (*amqp.Connection, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)

	client, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}

	return client, err
}

type MessageBroker struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewMessageBroker(connection *amqp.Connection) *MessageBroker {
	channel, err := connection.Channel()
	if err != nil {
		return nil
	}

	queue, err := channel.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil
	}

	return &MessageBroker{
		channel,
		queue,
	}
}

func (m *MessageBroker) PublishMessage(ctx context.Context, messageBody []byte, contentType string) error {
	err := m.channel.PublishWithContext(
		ctx,
		"",
		m.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        messageBody,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *MessageBroker) ConsumeMessages() <-chan Message {
	consume, err := m.channel.Consume(
		m.queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil
	}

	messages := make(chan Message)

	go func() {
		for msg := range consume {
			messages <- Message{
				Body:        msg.Body,
				ContentType: msg.ContentType,
			}
		}
	}()

	return messages
}
