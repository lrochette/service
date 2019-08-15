package model

// Author represents a row in the authors table
type Author struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	UUID      string `db:"uuid"`
}
