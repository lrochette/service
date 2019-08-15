package db

import (
	"context"

	"github.com/JPZ13/service/db/model"
)

// Posts holds db methods related to posts
type Posts interface {
	CreatePost(ctx context.Context, post *model.Post) error
	GetPost(ctx context.Context, postUUID string) (*model.Post, error)
}

type postsQueries struct {
	database
}

// CreatePost creates a row in the posts table
func (q *postsQueries) CreatePost(ctx context.Context, post *model.Post) error {
	_, err := q.db.NamedExec(`INSERT INTO posts
			(author_uuid, body, uuid)
			VALUES (:author_uuid, :body, :uuid)`,
		*post)

	return err
}

// GetPost gets a post row by the uuid
func (q *postsQueries) GetPost(ctx context.Context, postUUID string) (*model.Post, error) {
	post := &model.Post{}
	err := q.db.Get(post, `SELECT (author_uuid, body, uuid, timestamp)
			FROM posts WHERE uuid=$1`,
		postUUID)

	return post, err
}
