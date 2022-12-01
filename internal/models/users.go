package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
    VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, name, email, hashedPassword)
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}

		return err
	}

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
