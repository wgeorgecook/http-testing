package api

import (
	"errors"
	"github.com/wgeorgecook/testing-http/internal/pkg/utils/errs"
	"github.com/wgeorgecook/testing-http/internal/pkg/utils/mocks"
	"io"
	"net/http"
	"testing"
)

func TestClient_GetResourceByID(t *testing.T) {
	type test struct {
		name   string
		id     string
		body   string
		result int
		err    error
	}

	var tests = []test{
		{
			name:   "Get Resource With Valid ID Is Successful",
			id:     "1",
			result: http.StatusOK,
			body:   mocks.MockResourseId1,
			err:    nil,
		},
		{
			name:   "Get Resource With Empty ID Is Bad Request",
			id:     "",
			result: http.StatusBadRequest,
			err:    errs.ErrIdRequired,
			body:   errs.ErrIdRequired.Error(),
		},
		{
			name:   "Get Resource With ID == 3 Is Not Found",
			id:     "3",
			result: http.StatusNotFound,
			err:    errs.ErrIdNotFound,
			body:   errs.ErrIdNotFound.Error(),
		},
	}

	for _, tt := range tests {
		mockClient := &mocks.HTTPClientMock{}
		c := Client{HttpClient: mockClient}
		resp, err := c.GetResourceByID(tt.id)
		if !errors.Is(err, tt.err) {
			t.Errorf("GetResourceByID(%s) error = %v, wantErr '%v'", tt.id, err, tt.err)
		}
		if resp.StatusCode != tt.result {
			t.Errorf("GetResourceByID(%v) = %v, want %v", tt.id, resp.StatusCode, tt.result)
		}
		if resp.Body == nil && tt.body == "" {
			continue
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error reading response body: %v", err)
		}
		if string(bodyBytes) != tt.body {
			t.Errorf("GetResourceByID(%v) response = %v, want %v", tt.id, string(bodyBytes), tt.body)
		}
	}
}
