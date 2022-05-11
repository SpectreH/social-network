package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/internal/config"
	"social-network/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) SignIn(w http.ResponseWriter, r *http.Request) {
	var authDada models.Auth
	if err := json.NewDecoder(r.Body).Decode(&authDada); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := m.authenticate(authDada, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, _ := json.MarshalIndent(res, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) AuthMe(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(config.SESSION_NAME)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Println(c, err)

	user := models.User{Nickname: "Test"}

	out, _ := json.MarshalIndent(user, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) authenticate(authData models.Auth, w http.ResponseWriter) (models.FormValidationResponse, error) {
	id, err := m.DB.CheckEmailExistence(authData.Email)
	if err != nil {
		return models.FormValidationResponse{}, err
	}

	if id == 0 {
		return models.FormValidationResponse{
			OK:      false,
			Input:   "email",
			Message: "This email is not registered",
		}, nil
	}

	hash, err := m.DB.GetUserHash(id)
	if err != nil {
		return models.FormValidationResponse{}, err
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(authData.Password)) != nil {
		return models.FormValidationResponse{
			OK:      false,
			Input:   "password",
			Message: "Invalid password",
		}, nil
	}

	err = m.DB.UpdateSessionToken(createSessionToken(w), id)

	return models.FormValidationResponse{OK: true}, err
}
