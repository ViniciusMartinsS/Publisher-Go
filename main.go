package main

import (
	"log"

	"github.com/ViniciusMartinss/publisher-go/handler"
	"github.com/ViniciusMartinss/publisher-go/infrastructure/api"
	"github.com/ViniciusMartinss/publisher-go/infrastructure/broker"
	"github.com/ViniciusMartinss/publisher-go/usecase"
)

func main() {
	broker := broker.NewBrokerBuilder().
		Connect().
		ChannelSetup().
		BuildQueue("amqpGoQueue")

	publisherUsecase := usecase.NewPublsisherUsecase(broker)
	publisherHandler := handler.NewPublsisherHandler(publisherUsecase)

	server := api.NewServer(publisherHandler).
		Create()

	go server.Start()
	log.Printf("[INFO] Server Initialized Successfully. Running on: %s\n", "8090")

	server.Stop()
	broker.Disconnect()
}
