package core

import (
	"log"

	"github.com/jpz13/service/model"
)

type Config struct{}

type DummyService interface {
	DummyFunction(hoursList *[]model.DummyRequest) *[]model.DummyResponse
}

type dummyService struct {
	logger log.Logger
}

func NewDummyService(config *Config) DummyService {
	return &dummyService{}
}
