package server

import (
	"github.com/wgeorgecook/testing-http/internal/pkg/utils/errs"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getResourceHandler(t *testing.T) {
	type test struct {
		name           string
		inId           string
		wantBody       string
		wantStatusCode int
	}

	var tests = []test{
		{
			name:           "No ID is Bad Request",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       errs.ErrIdRequired.Error() + "\n",
		},
		{
			name:           "ID 1 is Returned",
			inId:           "1",
			wantStatusCode: http.StatusOK,
			wantBody:       `{"id":"1"}`,
		},
		{
			name:           "ID 3 is Not Found",
			inId:           "3",
			wantStatusCode: http.StatusNotFound,
			wantBody:       errs.ErrIdNotFound.Error() + "\n",
		},
	}

	var getResp = func(id string) (*http.Response, []byte) {
		req := httptest.NewRequest("GET", "http://example.com/api/v1/resources?id="+id, nil)
		w := httptest.NewRecorder()
		getResourceHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		return resp, body
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, body := getResp(tt.inId)
			if string(body) != tt.wantBody {
				t.Errorf("test %q body check: got %q, want %q", tt.name, string(body), tt.wantBody)
			}
			if resp.StatusCode != tt.wantStatusCode {
				t.Errorf("test %q status code check: got %d, want %d", tt.name, resp.StatusCode, tt.wantStatusCode)
			}
		})
	}
}
