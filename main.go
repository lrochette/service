package main

import (
	"bytes"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/docker/saas-mega/services/billing-api/config"
)

func main() {
	buf := bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)

	server, err := newServer(&config.Server{
		ListenPort: 8080,
		Logger:     logger,
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
