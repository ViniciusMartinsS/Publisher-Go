package usecase

import (
	"log"

	"github.com/ViniciusMartinss/publisher-go/domain"
)

type publisher struct {
	broker domain.Broker
}

func NewPublsisherUsecase(broker domain.Broker) domain.PublisherUsecase {
	return publisher{broker}
}

func (p publisher) Publish() bool {
	log.Println("Publisher Usecase was called successfully!")

	content := domain.Content{
		Type: "text/plain",
		Body: []byte("Testing second request to queue"),
	}

	published := p.broker.
		Publish("amqpGoQueue", content)

	return published
}
