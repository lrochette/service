package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/docker/saas-mega/services/billing-api/config"
)

type server struct {
	apiHandler http.Handler
	listenPort int
	serv       *http.Server
	config     *config.Server
	logger     log.Logger
}

func newServer(config *config.Server) (*server, error) {
	return &server{
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

	log.Println("Starting support analytics API server")

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

	s.logger.Print("Started support analytics API server")
}

func (s *server) stop() error {
	s.logger.Print("Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.serv.Shutdown(ctx)

	s.logger.Print("Server stopped")
	return err
}
