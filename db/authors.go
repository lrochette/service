package db

import (
	"context"

	"github.com/JPZ13/service/db/model"
)

// Authors holds db methods related to authors
type Authors interface {
	CreateAuthor(ctx context.Context, author *model.Author) error
	GetAuthor(ctx context.Context, authorID string) (*model.Author, error)
}

type authorsQueries struct {
	database
}

// CreateAuthor creates a row in the authors table
func (q *authorsQueries) CreateAuthor(ctx context.Context, author *model.Author) error {
	_, err := q.db.NamedExec(`INSERT INTO authors
			(first_name, last_name, uuid)
			VALUES (:first_name, :last_name, :uuid)`,
		*author)

	return err
}

// GetAuthor gets an author row by the uuid
func (q *authorsQueries) GetAuthor(ctx context.Context, authorID string) (*model.Author, error) {
	author := &model.Author{}
	err := q.db.Get(author, `SELECT first_name, last_name, uuid
			FROM authors WHERE uuid=$1`,
		authorID)
	return author, err
}
