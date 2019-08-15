package core

import (
	"context"

	"github.com/JPZ13/service/model"
	uuid "github.com/satori/go.uuid"
)

// AuthorsService holds methods related to authors
type AuthorsService interface {
	CreateAuthor(ctx context.Context, author *model.CreateAuthorRequest) (*model.CreateAuthorResponse, error)
}

type authorsService struct {
	baseService
}

// CreateAuthor creates an author
func (s *authorsService) CreateAuthor(ctx context.Context, author *model.CreateAuthorRequest) (*model.CreateAuthorResponse, error) {
	dbAuthor := translateCreateAuthorRequestToDBAuthor(author)

	// generate uuid
	dbAuthor.UUID = uuid.NewV4().String()

	// write author to db
	err := s.db.CreateAuthor(ctx, dbAuthor)
	if err != nil {
		return nil, err
	}

	response := translateDBAuthorToCreateAuthorResponse(dbAuthor)

	return response, nil
}
