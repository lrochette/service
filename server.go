package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/JPZ13/service/api"
	"github.com/JPZ13/service/core"
	"github.com/JPZ13/service/db"
)

type server struct {
	apiHandler http.Handler
	listenPort int
	serv       *http.Server
	config     *serverConfig
	logger     log.Logger
}

// serverConfig holds all config needed to start a support analytics service instance
type serverConfig struct {
	ListenPort int
	Logger     log.Logger
}

func newServer(config *serverConfig) (*server, error) {
	db, err := db.New(&db.Config{
		Driver:     *dbDriver,
		Username:   *dbUser,
		Password:   *dbPassword,
		Host:       *dbHost,
		DBName:     *dbName,
		SSLMode:    *dbSSLMode,
		MaxRetries: *dbMaxRetries,
		Port:       *dbPort,
	})
	if err != nil {
		return nil, err
	}

	dummyService := core.New(&core.Config{
		DB: db,
	})

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
