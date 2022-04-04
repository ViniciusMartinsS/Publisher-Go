package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ViniciusMartinss/publisher-go/domain"
)

type serverApi interface {
	Create() server
	Start()
	Stop()
}

type server struct {
	publishHandler domain.PublisherHandler
	server         *http.Server
}

func NewServer(publishHandler domain.PublisherHandler) serverApi {
	return server{publishHandler, nil}
}

func (s server) router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/publish", NewPublsisherApiHandler(s.publishHandler))

	return mux
}

func (s server) Create() server {
	s.server = &http.Server{
		Addr:    ":8090",
		Handler: s.router(),
	}
	return s
}

func (s server) Start() {
	s.server.ListenAndServe()
}

func (s server) Stop() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
	)

	<-ctx.Done()
	log.Println("[INFO] Server shutting down!")

	err := s.server.
		Shutdown(context.Background())

	if err != nil {
		log.Fatalf("[FAIL] Fail on stop the server, reason: %s\n", err.Error())
	}

	stop()
}
