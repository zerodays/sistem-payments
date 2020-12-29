package util

import (
	"github.com/gorilla/mux"
	"github.com/zerodays/sistem-payments/internal/handle/errors"
	"net/http"
	"strconv"
)

func IdFromRequest(w http.ResponseWriter, r *http.Request) (int, bool) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		errors.Response(w, errors.InvalidID)
		return 0, false
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errors.Response(w, errors.InvalidID)
		return 0, false
	}

	return id, true
}
