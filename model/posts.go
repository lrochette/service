package model

// CreatePostRequest is the request body to
// create a post
type CreatePostRequest struct {
	Body string `json:"body"`
}

// PostResponse is the response body when
// requesting a post
type PostResponse struct {
	AuthorUUID string `json:"authorUUID"`
	Timestamp  string `json:"timestamp"`
	Body       string `json:"body"`
	PostUUID   string `json:"postUUID"`
}
