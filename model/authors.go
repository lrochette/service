package model

// CreateAuthorRequest is the request body to
// create an author
type CreateAuthorRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// CreateAuthorResponse is the response body when
// creating an author
type CreateAuthorResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UUID      string `json:"uuid"`
}
