package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Returns an instance of the application struct containing
// mocked dependencies
func newTestApplication(t *testing.T) *application {
	return &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}
}

// Type witch embeds a httptest.Server instance
type testServer struct {
	*httptest.Server
}

// Initialzes and returns a new instance of custom testServer type.
func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)
	return &testServer{ts}
}

// Makes a GET request to a given url path using test server client, and
// returns status code, headers and body.
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}
