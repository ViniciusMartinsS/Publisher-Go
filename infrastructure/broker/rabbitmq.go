package broker

import (
	"log"

	"github.com/streadway/amqp"
)

type brokerSetting struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

type broker struct{ brokerSetting }

func NewBrokerBuilder() broker {
	return broker{}
}

func (b broker) Connect() broker {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	b.connection = connection

	return b
}

func (b broker) ChannelSetup() broker {
	channel, err := b.connection.Channel()
	if err != nil {
		panic(err)
	}

	b.channel = channel

	return b
}

func (b broker) BuildQueue(name string) broker {
	queue, err := b.channel.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	b.queue = queue
	return b
}

func (b broker) Publish(queue string, message string) bool {
	//@TODO - Receive content as parameter with a generic struct (ContentType, Body) that matches to the Publish of amqp
	content := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}

	err := b.channel.Publish("", queue, false, false, content)
	if err != nil {
		log.Printf("[ERROR] Error occurred while publishing message to queue: %s\n - error: %v", queue, err.Error())
		return false
	}

	return true
}

func (b broker) Disconnect() {
	b.channel.Close()
	b.connection.Close()
}
