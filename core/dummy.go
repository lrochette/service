package core

import (
	"github.com/jpz13/service/model"
)

func (s *dummyService) DummyFunction(list *[]model.DummyRequest) *[]model.DummyResponse {
	response := new([]model.DummyResponse)
	return response
}
