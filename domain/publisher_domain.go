package domain

type PublisherHandler interface {
	Publish() bool
}

type PublisherUsecase interface {
	Publish() bool
}
