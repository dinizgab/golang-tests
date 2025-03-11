package service

import "fmt"

type NotificationService interface {
	Publish(queue string, metadata []byte) error
}

type notificationServiceImpl struct {
	conn BrokerConnection
}

func NewNotificationService(conn BrokerConnection) NotificationService {
	return &notificationServiceImpl{
		conn,
	}
}

func (n *notificationServiceImpl) Publish(queue string, metadata []byte) error {
	ch, err := n.conn.Channel()
	if err != nil {
		return fmt.Errorf("NotificationService.Publish: error creating channel - %w", err)
	}
	defer ch.Close()

    q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("NotificationService.Publish: error declaring queue - %w", err)
	}

    err = ch.Publish(
        "",
        q.Name,
        false,
        false,
        metadata,
    )
    if err != nil {
        return fmt.Errorf("NotificationService.Publish: error publishing message - %w", err)
    }

	return nil
}
