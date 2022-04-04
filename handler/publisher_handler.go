package handler

import "github.com/ViniciusMartinss/publisher-go/domain"

type publisher struct {
	publisherUsecase domain.PublisherUsecase
}

func NewPublsisherHandler(publisherUsecase domain.PublisherUsecase) domain.PublisherHandler {
	return publisher{publisherUsecase}
}

func (p publisher) Publish() bool {
	return p.publisherUsecase.
		Publish()
}
