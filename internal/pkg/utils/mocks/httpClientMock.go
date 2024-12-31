package mocks

import (
	"github.com/wgeorgecook/testing-http/internal/pkg/utils/errs"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const MockResourseId1 = `{"id": "1"}`

// HTTPClientMock is the type that receives requests
// during the tests for an api.Client.
type HTTPClientMock struct{}

// Get implements the api.HttpClient interface and
// receives requests made by the api.Client when
// client.HttpClient.Get is called.
// Note: Specific behaviors expected from the
// remote endpoint MUST be updated or implemented
// in this function for the tests to remain valid.
func (c *HTTPClientMock) Get(endpoint string) (*http.Response, error) {
	fullUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	id := fullUrl.Query().Get("id")
	switch id {
	case "":
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(strings.NewReader(errs.ErrIdRequired.Error())),
		}, errs.ErrIdRequired
	case "1":
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(MockResourseId1)),
		}, nil
	case "3":
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       io.NopCloser(strings.NewReader(errs.ErrIdNotFound.Error())),
		}, errs.ErrIdNotFound
	}

	return &http.Response{StatusCode: http.StatusNotImplemented}, errs.ErrUnimplemented
}
