package handlers

import (
	"database/sql"
	"net/http"
	"social-network/internal/config"
	"social-network/internal/database"
	"social-network/internal/database/sqlite"

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

func CheckSession(w http.ResponseWriter, r *http.Request) (int, error) {
	c, err := r.Cookie(config.SESSION_NAME)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return 0, err
	}

	id, err := Repo.DB.CheckSessionExistence(c.Value)
	if err != nil {
		destroySession(w)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return 0, err
	}

	return id, nil
}

func destroySession(w http.ResponseWriter) {
	c := http.Cookie{
		Name:   "session_token",
		MaxAge: -1}
	http.SetCookie(w, &c)
}

// createSessionToken creates token for cookies and database
func createSessionToken(w http.ResponseWriter) string {
	sessionToken := uuid.NewV4().String()

	http.SetCookie(w, &http.Cookie{
		Name:     config.SESSION_NAME,
		Value:    sessionToken,
		MaxAge:   0,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	return sessionToken
}
