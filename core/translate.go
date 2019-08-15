package core

import (
	dbmodel "github.com/JPZ13/service/db/model"
	"github.com/JPZ13/service/model"
)

func translateCreateAuthorRequestToDBAuthor(author *model.CreateAuthorRequest) *dbmodel.Author {
	return &dbmodel.Author{
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}
}

func translateDBAuthorToAuthorResponse(dbAuthor *dbmodel.Author) *model.AuthorResponse {
	return &model.AuthorResponse{
		FirstName: dbAuthor.FirstName,
		LastName:  dbAuthor.LastName,
		UUID:      dbAuthor.UUID,
	}
}

func translateCreatePostRequestToDBPost(authorID string, post *model.CreatePostRequest) *dbmodel.Post {
	return &dbmodel.Post{
		AuthorUUID: authorID,
		Body:       post.Body,
	}
}

func translateDBPostToPostResponse(dbpost *dbmodel.Post) *model.PostResponse {
	return &model.PostResponse{
		AuthorUUID: dbpost.AuthorUUID,
		Body:       dbpost.Body,
		Timestamp:  dbpost.Timestamp,
		PostUUID:   dbpost.PostUUID,
	}
}
