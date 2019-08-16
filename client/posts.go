package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/JPZ13/service/model"
)

// Posts houses http methods related to blog posts
type Posts interface {
	CreatePost(ctx context.Context, authorUUID string, post *model.CreatePostRequest) (*model.PostResponse, error)
	GetPost(ctx context.Context, authorUUID, postUUID string) (*model.PostResponse, error)
}

type postsClient struct {
	client
}

// CreatePost is the client to create a post
func (c *postsClient) CreatePost(ctx context.Context, authorUUID string, post *model.CreatePostRequest) (*model.PostResponse, error) {
	url := c.serviceURI + v1API + fmt.Sprintf("/authors/%s/posts", authorUUID)

	req, err := formatJSONRequest(http.MethodPost, url, post)
	if err != nil {
		return nil, err
	}

	res := &model.PostResponse{}
	_, err = executeRequest(c.web, req, res)
	return res, err
}

// GetPost is the client to get a post
func (c *postsClient) GetPost(ctx context.Context, authorUUID, postUUID string) (*model.PostResponse, error) {
	url := c.serviceURI + v1API + fmt.Sprintf("/authors/%s/posts/%s", authorUUID, postUUID)

	req, err := formatJSONRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res := &model.PostResponse{}
	_, err = executeRequest(c.web, req, res)
	return res, err
}
