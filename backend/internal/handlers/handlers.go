package handlers

import (
	"database/sql"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Repository is the repository type (Repository pattern)
type Repository struct {
}

// CreateNewRepo creates a new repository
func CreateNewRepo(conn *sql.DB) *Repository {
	return &Repository{}
}

// Repo is the repository used by the handlers
var Repo *Repository

// SetNewHandlers sets the repository for the handlers
func SetNewHandlers(r *Repository) {
	Repo = r
}

// createSessionToken creates token for cookies and database
func createSessionToken(w http.ResponseWriter) string {
	sessionToken := uuid.NewV4().String()

	http.SetCookie(w, &http.Cookie{
		Name:     "sn_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(1200 * time.Second),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	return sessionToken
}
