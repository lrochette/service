package rest

import (
	"net/http"

	"github.com/JPZ13/service/model"
	restful "github.com/emicklei/go-restful"
)

type dummyResource struct {
	baseResource
}

// DummyFunction is a placeholder
func (r *dummyResource) DummyFunction(request *restful.Request, response *restful.Response) {
	dummyRequest := new([]model.DummyRequest)
	if !decodeRequest(request, response, dummyRequest) {
		return
	}

	res := r.service.DummyFunction(dummyRequest)

	response.WriteHeaderAndEntity(http.StatusCreated, res)
}
