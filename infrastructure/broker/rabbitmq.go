package broker

import (
	"log"

	"github.com/ViniciusMartinss/publisher-go/domain"
	"github.com/streadway/amqp"
)

type Broker interface {
	Connect() broker
	ChannelSetup() broker
	BuildQueue(name string) broker
	Publish(queue string, content domain.Content) bool
	Disconnect()
}

type brokerSetting struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

type broker struct{ brokerSetting }

func NewBrokerBuilder() Broker {
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

func (b broker) Publish(queue string, content domain.Content) bool {
	message := amqp.Publishing{
		ContentType: content.Type,
		Body:        content.Body,
	}

	err := b.channel.Publish("", queue, false, false, message)
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
