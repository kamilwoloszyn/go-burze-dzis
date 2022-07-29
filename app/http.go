package app

import "net/http"

// The HTTP interface, for mocking, tests ...
type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}
