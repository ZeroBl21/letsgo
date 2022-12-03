package models

import (
	"database/sql"
	"os"
	"testing"
)

const uri = "test_web:pass@/test_snippetbox?parseTime=true&multiStatements=true"

// Creates a new connection to the test database and returns
// the connection pool
func newTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", uri)
  if err != nil {
    t.Fatal(err)
  }

  script, err := os.ReadFile("./testdata/setup.sql")
  if err != nil {
    t.Fatal(err)
  }

  _, err = db.Exec(string(script))
  if err != nil {
    t.Fatal(err)
  }

  t.Cleanup(func() {
    script, err := os.ReadFile("./testdata/teardown.sql")
    if err != nil {
      t.Fatal(err)
    }

    _, err = db.Exec(string(script))
    if err != nil {
      t.Fatal(err)
    }
    
    db.Close()
  })

  return db
}
