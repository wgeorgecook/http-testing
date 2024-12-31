package api

import (
	"net/http"
)

// HttpClient is an interface that partially
// supports the http.Client interface and
// allows mocking requests made on it.
type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

// Client is the default way to interact with
// the API SDK. Requests to the remote endpoints
// are made via the HttpClient provided.
type Client struct {
	HttpClient HttpClient
}

const (
	baseUrl = "http://localhost:8080/api/v1/resources"
)

// GetResourceByID fetches the provided ID from the remote
// URL.
func (c *Client) GetResourceByID(id string) (*http.Response, error) {
	var url = baseUrl + "?id=" + id
	return c.HttpClient.Get(url)
}
