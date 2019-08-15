package main

import (
	"bytes"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ianschenck/envflag"
)

var (
	dbUser       = envflag.String("POSTGRES_USER", "postgres", "Username for db")
	dbPassword   = envflag.String("POSTGRES_PASSWORD", "development", "Password for db")
	dbName       = envflag.String("POSTGRES_DB", "postgres", "DB name")
	dbDriver     = envflag.String("DB_DRIVER", "postgres", "Database driver")
	dbHost       = envflag.String("DB_HOST", "postgres", "Database host")
	dbSSLMode    = envflag.String("DB_SSL_MODE", "disable", "Database ssl connection mode")
	dbMaxRetries = envflag.Int("DB_MAX_RETRIES", 10, "Max retries to connect to db")
	dbPort       = envflag.Int("DB_PORT", 5432, "Database port")
)

func main() {
	envflag.Parse()

	buf := bytes.Buffer{}
	logger := log.New(&buf, "logger: ", log.Lshortfile)

	server, err := newServer(&serverConfig{
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
