package handlers

import (
	"net/http"
	"social-network/internal/config"
)

func (m *Repository) LogOut(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   config.SESSION_NAME,
		MaxAge: -1}
	http.SetCookie(w, &c)

	return
}
