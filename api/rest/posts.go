package rest

import (
	"context"
	"net/http"

	"github.com/JPZ13/service/model"
	restful "github.com/emicklei/go-restful"
)

type postsResource struct {
	baseResource
}

// CreatePost is the decoding/error layer to create a post
func (r *postsResource) CreatePost(request *restful.Request, response *restful.Response) {
	authorID := request.PathParameter("author-id")

	createPostRequest := &model.CreatePostRequest{}
	if !decodeRequest(request, response, createPostRequest) {
		return
	}

	ctx := context.Background()
	res, err := r.service.CreatePost(ctx, authorID, createPostRequest)
	if err != nil {
		encodeErrorWithStatus(response, err, http.StatusBadRequest)
	}

	response.WriteHeaderAndEntity(http.StatusCreated, res)
}
