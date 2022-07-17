package app

import "net/http"

type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}
