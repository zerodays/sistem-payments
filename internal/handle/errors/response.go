package errors

import (
	"encoding/json"
	"net/http"
)

// Response writes error `err` to response.
func Response(w http.ResponseWriter, err ResponseError) {
	// Create JSON body.
	res, _ := json.Marshal(map[string]interface{}{
		"error_code": err.Code,
		"metadata":   err.Metadata,
	})

	// Set headers and write response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	_, _ = w.Write(res)
}
