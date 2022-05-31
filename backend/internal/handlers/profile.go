package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"social-network/internal/models"
	"strconv"
)

func (m *Repository) ProfileRequestToFollow(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	queries := r.URL.Query()
	id := queries.Get("id")
	if id == "" {
		http.Error(w, errors.New("No id inside query").Error(), http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "Profile with this id doesn't exist!",
		}
	}

	if intId == uid {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "You can't follow yourself!",
		}
	}

	if response.OK {
		res, err := m.DB.CheckFollowRequest(uid, intId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if res != 0 {
			response = models.FormValidationResponse{
				OK:      false,
				Message: "Follow request already sended!",
			}
		}
	}

	if response.OK {
		err = m.DB.InsertUserFollowRequest(uid, intId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Message = "Follow request successfully sended!"
	}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}

func (m *Repository) ProfileUnFollow(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	queries := r.URL.Query()
	id := queries.Get("id")
	if id == "" {
		http.Error(w, errors.New("No id inside query").Error(), http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "Profile with this id doesn't exist!",
		}
	}

	if intId == uid {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "You can't unfollow yourself!",
		}
	}

	if response.OK {
		err = m.DB.UnFollow(uid, intId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Message = "You stopped following this profile!"
	}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}

func (m *Repository) ProfileFollow(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	queries := r.URL.Query()
	id := queries.Get("id")
	if id == "" {
		http.Error(w, errors.New("No id inside query").Error(), http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "Profile with this id doesn't exist!",
		}
	}

	if intId == uid {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "You can't follow yourself!",
		}
	}

	if response.OK {
		private, err := m.DB.CheckProfileIsPivate(intId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if private {
			response = models.FormValidationResponse{
				OK:      false,
				Message: "You must request to follow this profile!",
			}
		}
	}

	if response.OK {
		res, err := m.DB.CheckAlreadyFollowed(uid, intId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if res != 0 {
			response = models.FormValidationResponse{
				OK:      false,
				Message: "You already following that user!",
			}
		}
	}

	if response.OK {
		err = m.DB.FollowUser(uid, intId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Message = "You now following this profile!"
	}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}

func (m *Repository) GetProfile(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	queries := r.URL.Query()
	id := queries.Get("id")
	if id == "" {
		http.Error(w, errors.New("No id inside query").Error(), http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "Profile with this id doesn't exist!",
		}
	}

	profile, err := m.DB.GetUserProfile(intId)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "Profile with this id doesn't exist!",
		}
	}

	if intId == uid {
		profile.IsMyProfile = true
	}

	if !response.OK {
		out, _ := json.MarshalIndent(response, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	res, err := m.DB.CheckAlreadyFollowed(uid, intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if res == 1 {
		profile.Following = true
	}

	out, _ := json.MarshalIndent(profile, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
