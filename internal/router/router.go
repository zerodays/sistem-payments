package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

// addRoute adds route to router.
func addRoute(router *mux.Router, r Route) {
	// Add handlers for specified methods.
	if r.GET != nil {
		router.Handle(r.Path, addMiddleware(r.GET, r.AuthorizedOnly, r.CustomContentType)).
			Methods("GET").
			Name(r.Name)
	}

	if r.PUT != nil {
		router.Handle(r.Path, addMiddleware(r.PUT, r.AuthorizedOnly, r.CustomContentType)).
			Methods("PUT").
			Name(r.Name)
	}

	if r.POST != nil {
		router.Handle(r.Path, addMiddleware(r.POST, r.AuthorizedOnly, r.CustomContentType)).
			Methods("POST").
			Name(r.Name)
	}

	if r.DELETE != nil {
		router.Handle(r.Path, addMiddleware(r.DELETE, r.AuthorizedOnly, r.CustomContentType)).
			Methods("DELETE").
			Name(r.Name)
	}

	// Add CORS handle for OPTIONS request
	router.Handle(r.Path, addMiddleware(http.HandlerFunc(corsHandler), false, false)).
		Methods("OPTIONS")
}

// NewRouter creates new router and populates it with routes.
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Populate router with API routes.
	s := router.PathPrefix("/api/v1").Subrouter()
	for _, r := range apiRoutes() {
		addRoute(s, r)
	}

	return router
}
