package models

import (
	"database/sql"
	"time"
)

// User model align with te database "users" table.
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// Wraps database connection pool.
type UserModel struct {
	DB *sql.DB
}

// Insert method add a new record to the "users" table.
func (m *UserModel) Insert(name, email, password string) error {
	// TODO: Insert a User to the database
	return nil
}

// Verify whether a user exists with the provided email address and password
// This returns the relevant user ID if they do
func (m *UserModel) Authenticate(email, password string) (int, error) {
	// TODO: Check in the database
	return 0, nil
}

// Check if the user exists with a specific ID
func (m *UserModel) Exists(id int) (bool, error) {
	// TODO
	return false, nil
}
