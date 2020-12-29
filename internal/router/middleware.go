package router

import (
	"github.com/google/uuid"
	"github.com/justinas/alice"
	"github.com/rs/zerolog/hlog"
	"github.com/zerodays/sistem-auth/middleware"
	"github.com/zerodays/sistem-payments/internal/logger"
	"net/http"
	"time"
)

// createLoggerMiddleware creates logger middleware.
func createLoggerMiddleware() func(handler http.Handler) http.Handler {
	return hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		u := middleware.UserFromRequest(r)

		var userID string
		if u != nil {
			userID = u.UID
		}

		// Set request ID header if needed
		requestId := r.Header.Get("X-Request-ID")
		if requestId == "" {
			requestId = uuid.New().String()
			r.Header.Set("X-Request-ID", requestId)
		}

		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Bool("authenticated", u != nil).
			Str("user_id", userID).
			Str("request_id", requestId).
			Send()
	})
}

// writeCorsHeader writes CORS header to allow all connections.
func writeCorsHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization")
}

// corsHandler handles OPTIONS requests. Writes CORS header fields
// and responds with empty body and status 200.
func corsHandler(w http.ResponseWriter, _ *http.Request) {
	writeCorsHeader(w)

	_, _ = w.Write([]byte{})
}

// corsMiddleware writes CORS headers before further processing the request.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeCorsHeader(w)

		next.ServeHTTP(w, r)
	})
}

// responseTypeHeaderMiddleware sets Content-Type header to "application/json" before further processing the request.
func responseTypeHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

// addMiddleware adds necessary middleware to the handle.
func addMiddleware(handler http.Handler, authorizedOnly, customContentType bool) http.Handler {
	c := alice.New(hlog.NewHandler(logger.Log), middleware.Middleware, createLoggerMiddleware())

	if !customContentType {
		c = c.Append(responseTypeHeaderMiddleware)
	}

	c = c.Append(corsMiddleware)

	// TODO: Uncomment to enable user auth
	//if authorizedOnly {
	//	c = c.Append(middleware.RequiredMiddleware)
	//}

	return c.Then(handler)
}
