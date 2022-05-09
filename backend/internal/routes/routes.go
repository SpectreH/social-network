package routes

import (
	"net/http"
	"social-network/internal/handlers"
)

// SetRoutes sets handler and load server files
func SetRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/test", handlers.Repo.Test)

	return mux
}
