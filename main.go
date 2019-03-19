package main

import (
	"bytes"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/jpz13/service/api"
	"github.com/jpz13/service/config"
	"github.com/jpz13/service/core"
)

func main() {
	buf := bytes.Buffer{}
	logger := log.New(&buf, "logger: ", log.Lshortfile)

	server, err := newServer(&config.Server{
		ListenPort: 8080,
		Logger:     *logger,
	})

	if err != nil {
		log.Println("failed to init server")
	}

	server.start()

	// Wait for signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh
	log.Println("Caught signal:", sig)

	err = server.stop()
	if err != nil {
		log.Println("Error while shutting down server:", err)
	}
}

type server struct {
	apiHandler http.Handler
	listenPort int
	serv       *http.Server
	config     *config.Server
	logger     log.Logger
}

func newServer(config *config.Server) (*server, error) {
	dummyService := core.NewDummyService(&core.Config{})

	apiHandler := api.New(&api.Config{
		DummyService: dummyService,
		Logger:       config.Logger,
	})

	return &server{
		apiHandler: apiHandler,
		listenPort: config.ListenPort,
		logger:     config.Logger,
		config:     config,
	}, nil
}

func (s *server) start() {
	addr := net.JoinHostPort("0.0.0.0", strconv.Itoa(s.listenPort))

	serv := &http.Server{
		Addr:         addr,
		Handler:      s.apiHandler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("Starting API server")

	go func() {
		err := serv.ListenAndServe()
		// expected error on shutdown
		if err == http.ErrServerClosed {
			log.Printf("ListenAndServe response: %v", err)
		} else {
			log.Fatal("API server listen failed")
		}
	}()

	s.serv = serv

	s.logger.Print("Started API server")
}

func (s *server) stop() error {
	s.logger.Print("Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.serv.Shutdown(ctx)

	s.logger.Print("Server stopped")
	return err
}
