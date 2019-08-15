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

func translateDBAuthorToCreateAuthorResponse(dbAuthor *dbmodel.Author) *model.CreateAuthorResponse {
	return &model.CreateAuthorResponse{
		FirstName: dbAuthor.FirstName,
		LastName:  dbAuthor.LastName,
		UUID:      dbAuthor.UUID,
	}
}
