package core

import (
	"log"

	"github.com/JPZ13/service/db"
)

// Config holds core service configuration
type Config struct {
	DB db.Database
}

// Service holds core methods
type Service interface {
	AuthorsService
	PostsService
}

type coreService struct {
	authorsService
	postsService
}

type baseService struct {
	logger log.Logger
	db     db.Database
}

// New inits the core package
func New(config *Config) Service {
	base := baseService{
		db: config.DB,
	}

	return &coreService{
		authorsService{base},
		postsService{base},
	}
}
