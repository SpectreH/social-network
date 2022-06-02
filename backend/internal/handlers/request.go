package handlers

import (
	"encoding/json"
	"net/http"
	"social-network/internal/models"
)

func (m *Repository) AcceptRequest(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{OK: true}

	sourceId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	res, err := m.DB.CheckAlreadyUserFollowed(sourceId, uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if res != 0 {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "This user already following you!",
		}
	}

	if response.OK {
		err = m.DB.RemoveFollowRequest(sourceId, uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = m.DB.FollowUser(sourceId, uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Message = "Follow request accepted!"
	}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}

func (m *Repository) DeclineRequest(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{OK: true}

	sourceId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	if response.OK {
		err = m.DB.RemoveFollowRequest(sourceId, uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Message = "Follow request declined!"
	}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}
