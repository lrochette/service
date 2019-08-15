package db

import (
	"context"

	"github.com/JPZ13/service/db/model"
)

// Authors holds db methods related to authors
type Authors interface {
	CreateAuthor(ctx context.Context, author *model.Author) error
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
