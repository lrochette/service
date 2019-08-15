package core

import (
	"log"

	"github.com/JPZ13/service/db"
	"github.com/JPZ13/service/model"
)

// Config holds core service configuration
type Config struct {
	DB db.Database
}

// DummyService holds core methods
type DummyService interface {
	DummyFunction(hoursList *[]model.DummyRequest) *[]model.DummyResponse
}

type dummyService struct {
	logger log.Logger
	db     db.Database
}

// NewDummyService inits the core package
func NewDummyService(config *Config) DummyService {
	return &dummyService{
		db: config.DB,
	}
}
