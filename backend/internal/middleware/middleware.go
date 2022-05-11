package mw

import (
	"database/sql"
	"net/http"
)

// Repository is the repository type (Repository pattern)
type Repository struct {
}

// CreateNewRepo creates a new repository
func CreateNewRepo(conn *sql.DB) *Repository {
	return &Repository{}
}

// Repo is the repository used by the middleware
var Repo *Repository

// SetNewMiddleware sets the repository for the middleware
func SetNewMiddleware(r *Repository) {
	Repo = r
}

func (m *Repository) SetupCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Credentials", "true")
		(w).Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")
		next.ServeHTTP(w, r)
	})
}
