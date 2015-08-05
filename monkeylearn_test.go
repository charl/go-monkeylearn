package gomonkeylearn

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the go-monkeylearn client being tested.
	client *Client

	// server is a test HTTPS server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTPS server along with a github.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup(apiToken string) {
	// Test server.
	mux = http.NewServeMux()
	server = httptest.NewTLSServer(mux)

	// go-monkeylearn client configured to use test server.
	client = NewClient(apiToken)
	u, _ := url.Parse(server.URL)
	client.BaseURL = u.String()
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

// testMethod ensures we only get the HTTP method we're looking for.
func testMethod(t *testing.T, r *http.Request, want string) {
	if want != r.Method {
		t.Errorf("Request method = %v, want %v", r.Method, want)
	}
}
