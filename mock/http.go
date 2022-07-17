package mock

import (
	"errors"
	"net/http"
)

var (
	ErrNotMocked = errors.New("not mocked")
)

type MockedHTTP struct {
	MockedDoFunc func(*http.Request) (*http.Response, error)
}

func (m MockedHTTP) Do(req *http.Request) (*http.Response, error) {
	if m.MockedDoFunc != nil {
		return m.MockedDoFunc(req)
	}
	return nil, ErrNotMocked
}
