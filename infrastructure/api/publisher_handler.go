package api

import (
	"fmt"
	"net/http"

	"github.com/ViniciusMartinss/publisher-go/domain"
)

type publisherApiHandlers func(_ publisher, rw http.ResponseWriter, r *http.Request)

type publisher struct {
	publishHandler domain.PublisherHandler
}

func NewPublsisherApiHandler(publishHandler domain.PublisherHandler) publisher {
	return publisher{publishHandler}
}

var handleRequest = map[string]publisherApiHandlers{
	"POST": publisher.postHandler,
}

func (p publisher) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	requestHandler, exist := handleRequest[r.Method]
	if !exist {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.Header().
		Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	requestHandler(p, rw, r)
}

func (p publisher) postHandler(rw http.ResponseWriter, r *http.Request) {
	result := p.publishHandler.
		Publish()

	response := fmt.Sprintf("Welcome to the \"Just Enough Go\" blog series!! %v", result)
	rw.Write([]byte(response))
}
