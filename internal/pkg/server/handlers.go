package server

import (
	"encoding/json"
	"github.com/wgeorgecook/testing-http/internal/pkg/resources"
	"github.com/wgeorgecook/testing-http/internal/pkg/utils/errs"
	"net/http"
	"strconv"
)

func getResourceHandler(w http.ResponseWriter, req *http.Request) {
	idCheck, ok := req.URL.Query()["id"]
	if !ok || idCheck[0] == "" {
		http.Error(w, errs.ErrIdRequired.Error(), http.StatusBadRequest)
		return
	}

	id := idCheck[0]
	if id == "3" {
		http.Error(w, errs.ErrIdNotFound.Error(), http.StatusNotFound)
		return
	}

	r := resources.Resource{ID: id}
	rBytes, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(rBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(written))
	return
}
