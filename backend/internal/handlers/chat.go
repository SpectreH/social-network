package handlers

import (
	"encoding/json"
	"net/http"
	"social-network/internal/models"
	"strconv"
	"time"
)

func (m *Repository) GetChats(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	chats, err := m.DB.GetAllChats(uid, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groupChats, err := m.DB.GetAllChats(uid, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chats = append(chats, groupChats...)

	out, _ := json.MarshalIndent(chats, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) NewMessage(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cid, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := models.ChatMessage{
		ChatId:    cid,
		AuthorId:  uid,
		Content:   r.Form.Get("message"),
		CreatedAt: time.Now(),
	}

	err = m.DB.InsertChatMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FormValidationResponse{OK: true}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
