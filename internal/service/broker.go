package service

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type BrokerConnection interface {
	Channel() (BrokerChannel, error)
	Close() error
}

type BrokerChannel interface {
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args map[string]interface{}) (<-chan amqp.Delivery, error)
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args map[string]interface{}) (amqp.Queue, error)
	Publish(exchange, key string, mandatory, immediate bool, msg []byte) error
	Close() error
}

type brokerConnectionImpl struct {
	conn *amqp.Connection
}

type brokerChannelImpl struct {
	ch *amqp.Channel
}

func NewBrokerConnection() (BrokerConnection, error) {
	dsn := os.Getenv("BROKER_DSN")

	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, fmt.Errorf("NewBrokerConnection: error dialing connection - %w", err)
	}

	return &brokerConnectionImpl{conn}, nil
}

func (bc *brokerConnectionImpl) Channel() (BrokerChannel, error) {
	ch, err := bc.conn.Channel()
	if err != nil {
		return nil, err
	}

	return &brokerChannelImpl{ch}, nil
}

func (bc *brokerConnectionImpl) Close() error {
	return bc.conn.Close()
}

func (bc *brokerChannelImpl) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args map[string]interface{}) (<-chan amqp.Delivery, error) {
	return bc.ch.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

func (bc *brokerChannelImpl) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args map[string]interface{}) (amqp.Queue, error) {
	return bc.ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}

func (bc *brokerChannelImpl) Publish(exchange, key string, mandatory, immediate bool, msg []byte) error {
	return bc.ch.Publish(
		exchange,
		key,
		mandatory,
		immediate,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)
}

func (bc *brokerChannelImpl) Close() error {
	return bc.ch.Close()
}
