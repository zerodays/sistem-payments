package util

import (
	"encoding/json"
	"github.com/zerodays/sistem-payments/internal/handle/errors"
	"github.com/zerodays/sistem-payments/internal/logger"
	"net/http"
)

// ParseJSON parses json from request body into destination.
// Errors are written to response writer. Return weather
// parsing was successful.
func ParseJSON(w http.ResponseWriter, r *http.Request, dest interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dest)
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.InvalidJSON)

		return false
	}

	return true
}
