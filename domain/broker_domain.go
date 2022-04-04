package domain

type Broker interface {
	Publish(queue string, message Content) bool
}

type Content struct {
	Type string
	Body []byte
}
