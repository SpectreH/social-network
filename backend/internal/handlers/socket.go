package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// CreateSocketReader is the handler for new web socket connection
func (m *Repository) CreateSocketReader(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
		r.Body.Close()
	}()

	fullName, err := m.DB.GetUserFullName(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	m.SR.AppendNewConnection(con, fullName, uid)
}

func (m *Repository) GetRequestList(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	req, err := m.DB.GetUserFollowRequests(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	groupReq, err := m.DB.GetGroupFollowRequests(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req = append(req, groupReq...)

	out, _ := json.MarshalIndent(req, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
