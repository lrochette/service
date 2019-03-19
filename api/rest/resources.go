package rest

import (
	"log"

	"github.com/jpz13/service/core"
)

type baseResource struct {
	logger  log.Logger
	service core.DummyService
}

type Resources struct {
	dummyResource
}

type Config struct {
	DummyService core.DummyService
	Logger       log.Logger
}

func NewResources(config *Config) Resources {
	br := baseResource{
		logger:  config.Logger,
		service: config.DummyService,
	}

	return Resources{
		dummyResource{br},
	}
}
