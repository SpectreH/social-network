package handlers

import (
	"database/sql"
	"net/http"
	"social-network/internal/config"
	"social-network/internal/database"
	"social-network/internal/database/sqlite"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Repository is the repository type (Repository pattern)
type Repository struct {
	DB database.DatabaseRepo
}

// CreateNewRepo creates a new repository
func CreateNewRepo(conn *sql.DB) *Repository {
	return &Repository{
		DB: sqlite.SetSqliteRepo(conn),
	}
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
		Name:     config.SESSION_NAME,
		Value:    sessionToken,
		Expires:  time.Now().Add(config.SESSION_EXPIRATION_TIME),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	return sessionToken
}
