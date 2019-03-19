package config

import (
	"log"
)

// Server holds all config needed to start a support analytics service instance
type Server struct {
	ListenPort int
	Logger     log.Logger
}
