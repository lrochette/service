package rest

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/jpz13/service/model"
)

type dummyResource struct {
	baseResource
}

func (r *dummyResource) DummyFunction(request *restful.Request, response *restful.Response) {
	dummyRequest := new([]model.DummyRequest)
	if !decodeRequest(request, response, dummyRequest) {
		return
	}

	res := r.service.DummyFunction(dummyRequest)

	response.WriteHeaderAndEntity(http.StatusCreated, res)
}
