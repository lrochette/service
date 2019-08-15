package model

// Post represents a row in the posts table
type Post struct {
	AuthorUUID string `db:"author_uuid"`
	Timestamp  string `db:"timestamp"`
	Body       string `db:"body"`
	PostUUID   string `db:"uuid"`
}
