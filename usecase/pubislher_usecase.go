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

	published := p.broker.
		Publish("amqpGoQueue", "Hello From UseCase")

	return published
}
