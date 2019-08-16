package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/JPZ13/service/model"
)

// Authors holds http authors methods
type Authors interface {
	CreateAuthor(ctx context.Context, author *model.CreateAuthorRequest) (*model.AuthorResponse, error)
	GetAuthor(ctx context.Context, authorUUID string) (*model.AuthorResponse, error)
}

type authorsClient struct {
	client
}

// CreateAuthor is the http client to create an author
func (c *authorsClient) CreateAuthor(ctx context.Context, author *model.CreateAuthorRequest) (*model.AuthorResponse, error) {
	url := c.serviceURI + v1API + "/authors"

	req, err := formatJSONRequest(http.MethodPost, url, author)
	if err != nil {
		return nil, err
	}

	res := &model.AuthorResponse{}
	_, err = executeRequest(c.web, req, res)
	return res, err
}

// GetAuthor is the http client to get an author by id
func (c *authorsClient) GetAuthor(ctx context.Context, authorUUID string) (*model.AuthorResponse, error) {
	url := c.serviceURI + v1API + fmt.Sprintf("/authors/%s", authorUUID)

	req, err := formatJSONRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res := &model.AuthorResponse{}
	_, err = executeRequest(c.web, req, res)
	return res, err
}
