package core

import (
	"context"

	"github.com/JPZ13/service/model"
	uuid "github.com/satori/go.uuid"
)

// PostsService holds methods related to authors
type PostsService interface {
	CreatePost(ctx context.Context, authorID string, post *model.CreatePostRequest) (*model.PostResponse, error)
	GetPost(ctx context.Context, authorUUID, postUUID string) (*model.PostResponse, error)
}

type postsService struct {
	baseService
}

// CreatePost creates a post
func (s *authorsService) CreatePost(ctx context.Context, authorID string, post *model.CreatePostRequest) (*model.PostResponse, error) {
	dbPost := translateCreatePostRequestToDBPost(authorID, post)

	// generate uuid
	dbPost.PostUUID = uuid.NewV4().String()

	// write post to db
	err := s.db.CreatePost(ctx, dbPost)
	if err != nil {
		return nil, err
	}

	// get post from db to populate response
	dbPost, err = s.db.GetPost(ctx, authorID, dbPost.PostUUID)
	if err != nil {
		return nil, err
	}

	response := translateDBPostToPostResponse(dbPost)

	return response, nil
}

// GetPost gets a post from the db based on ids
func (s *authorsService) GetPost(ctx context.Context, authorUUID, postUUID string) (*model.PostResponse, error) {
	// get post from db to populate response
	dbPost, err := s.db.GetPost(ctx, authorUUID, postUUID)
	if err != nil {
		return nil, err
	}

	response := translateDBPostToPostResponse(dbPost)

	return response, nil
}
