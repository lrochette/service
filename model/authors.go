package model

// CreateAuthorRequest is the request body to
// create an author
type CreateAuthorRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// AuthorResponse is the response body when
// requesting an author
type AuthorResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UUID      string `json:"uuid"`
}
