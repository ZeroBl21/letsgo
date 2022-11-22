package models

import (
	"database/sql"
	"time"
)

// Holds data for an individual Snippet.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new Snippet into the database
// TODO
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 0, nil
}

// Return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Return top 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
