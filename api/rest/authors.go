package rest

import (
	"context"
	"net/http"

	"github.com/JPZ13/service/model"
	restful "github.com/emicklei/go-restful"
)

type authorsResource struct {
	baseResource
}

// CreateAuthor is the decoding/error layer to create an author
func (r *authorsResource) CreateAuthor(request *restful.Request, response *restful.Response) {
	createAuthorRequest := &model.CreateAuthorRequest{}
	if !decodeRequest(request, response, createAuthorRequest) {
		return
	}

	ctx := context.Background()
	res, err := r.service.CreateAuthor(ctx, createAuthorRequest)
	if err != nil {
		encodeErrorWithStatus(response, err, http.StatusBadRequest)
	}

	response.WriteHeaderAndEntity(http.StatusCreated, res)
}

// GetAuthor is the decoding/error layer to create an author
func (r *authorsResource) GetAuthor(request *restful.Request, response *restful.Response) {
	authorID := request.PathParameter("author-id")
	ctx := context.Background()

	res, err := r.service.GetAuthor(ctx, authorID)
	if err != nil {
		encodeErrorWithStatus(response, err, http.StatusBadRequest)
	}

	response.WriteHeaderAndEntity(http.StatusOK, res)
}
