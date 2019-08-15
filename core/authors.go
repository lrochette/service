package core

import (
	"context"

	"github.com/JPZ13/service/model"
	uuid "github.com/satori/go.uuid"
)

// AuthorsService holds methods related to authors
type AuthorsService interface {
	CreateAuthor(ctx context.Context, author *model.CreateAuthorRequest) (*model.AuthorResponse, error)
	GetAuthor(ctx context.Context, authorID string) (*model.AuthorResponse, error)
}

type authorsService struct {
	baseService
}

// CreateAuthor creates an author
func (s *authorsService) CreateAuthor(ctx context.Context, author *model.CreateAuthorRequest) (*model.AuthorResponse, error) {
	dbAuthor := translateCreateAuthorRequestToDBAuthor(author)

	// generate uuid
	dbAuthor.UUID = uuid.NewV4().String()

	// write author to db
	err := s.db.CreateAuthor(ctx, dbAuthor)
	if err != nil {
		return nil, err
	}

	response := translateDBAuthorToAuthorResponse(dbAuthor)

	return response, nil
}

// GetAuthor gets an author from the db and translates it to a response
func (s *authorsService) GetAuthor(ctx context.Context, authorID string) (*model.AuthorResponse, error) {
	dbAuthor, err := s.db.GetAuthor(ctx, authorID)
	if err != nil {
		return nil, err
	}

	response := translateDBAuthorToAuthorResponse(dbAuthor)

	return response, nil
}
