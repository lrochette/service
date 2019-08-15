package rest

import (
	"log"

	"github.com/JPZ13/service/core"
)

type baseResource struct {
	logger  log.Logger
	service core.DummyService
}

// Resources holds api resources
type Resources struct {
	dummyResource
}

// Config holds api resource config
type Config struct {
	DummyService core.DummyService
	Logger       log.Logger
}

// NewResources inits api resources
func NewResources(config *Config) Resources {
	br := baseResource{
		logger:  config.Logger,
		service: config.DummyService,
	}

	return Resources{
		dummyResource{br},
	}
}
