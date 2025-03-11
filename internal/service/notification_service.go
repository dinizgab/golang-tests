package service

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type NotificationService interface {
	Publish(metadata []byte) error
}

type notificationServiceImpl struct {
	conn    BrokerConnection
	channel BrokerChannel
	queue   amqp.Queue
}

func NewNotificationService(queue string, conn BrokerConnection) (NotificationService, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("NotificationService.NewNotificationService: error creating channel - %w", err)
	}

	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("NotificationService.NewNotificationService: error declaring queue - %w", err)
	}

	return &notificationServiceImpl{
		conn,
		ch,
		q,
	}, err
}

func (n *notificationServiceImpl) Publish(metadata []byte) error {
	err := n.channel.Publish(
		"",
		n.queue.Name,
		false,
		false,
		metadata,
	)
	if err != nil {
		return fmt.Errorf("NotificationService.Publish: error publishing message - %w", err)
	}

	return nil
}
