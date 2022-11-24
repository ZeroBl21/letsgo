package models

import (
	"database/sql"
	"errors"
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
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
  VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

// Return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
  stmt := `SELECT id, title, content, created, expires FROM snippets
  WHERE expires > UTC_TIMESTAMP() AND id = ?`
  s := &Snippet{}

  err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, ErrNoRecord
    }

    return nil, err
  }

	return s, nil
}

// Return top 10 most recently created snippets
// TODO
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
