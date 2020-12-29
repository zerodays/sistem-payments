package errors

import "net/http"

type ResponseError struct {
	StatusCode int
	Code       string
	Metadata   map[string]interface{}
}

var (
	InvalidJSON = ResponseError{
		StatusCode: http.StatusBadRequest,
		Code:       "invalid_json",
	}

	InvalidCredentials = ResponseError{
		StatusCode: http.StatusBadRequest,
		Code:       "invalid_credentials",
	}

	InsufficientPermissions = ResponseError{
		StatusCode: http.StatusForbidden,
		Code:       "insufficient_permission",
	}

	DatabaseError = ResponseError{
		StatusCode: http.StatusInternalServerError,
		Code:       "database_error",
	}

	InvalidID = ResponseError{
		StatusCode: http.StatusBadRequest,
		Code:       "invalid_id",
	}

	UnknownError = ResponseError{
		StatusCode: http.StatusInternalServerError,
		Code:       "unknown_error",
	}
)
