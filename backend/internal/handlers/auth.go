package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/internal/config"
	"social-network/internal/models"
)

func (m *Repository) SignIn(w http.ResponseWriter, r *http.Request) {
	createSessionToken(w)
}

func (m *Repository) AuthMe(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(config.SESSION_NAME)
	fmt.Println(c, err)

	user := models.User{Nickname: "Test"}

	out, _ := json.MarshalIndent(user, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
