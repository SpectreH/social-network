package routes

import (
	"net/http"
	"social-network/internal/handlers"

	middleware "social-network/internal/middleware"
)

// SetRoutes sets handler and load server files
func SetRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/signin", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.SignIn)))
	mux.Handle("/api/authme", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.AuthMe)))

	return mux
}
