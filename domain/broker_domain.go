package domain

type Broker interface {
	Publish(queue string, message string) bool
}
