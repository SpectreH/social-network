package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/internal/config"
	"social-network/internal/models"
	img "social-network/pkg/image-storage"
	"strconv"
	"time"
)

func (m *Repository) EventAcceptRequest(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{OK: true}

	eventId, err := getIdFromQuery(r, "id")
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	if response.OK {
		res, err := m.DB.CheckAlreadyParticipating(eventId, uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if res != 0 {
			response = models.FormValidationResponse{
				OK:      false,
				Message: "You are already chose option!",
			}
		}
	}

	if response.OK {
		err = m.DB.RemoveEventParticipatingRequest(eventId, uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = m.DB.InsertEventParticipater(eventId, uid, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Message = "Event request accepted!"
	}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}

func (m *Repository) EventDeclineRequest(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{OK: true}

	eventId, err := getIdFromQuery(r, "id")
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	if response.OK {
		res, err := m.DB.CheckAlreadyParticipating(eventId, uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if res != 0 {
			response = models.FormValidationResponse{
				OK:      false,
				Message: "You are already chose option!",
			}
		}
	}

	if response.OK {
		err = m.DB.RemoveEventParticipatingRequest(eventId, uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = m.DB.InsertEventParticipater(eventId, uid, false)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Message = "Event request declined!"
	}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) NewEvent(w http.ResponseWriter, r *http.Request) {
	_, err := CheckSession(w, r)
	if err != nil {
		return
	}

	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	gid, err := strconv.Atoi(r.Form.Get("groupId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	event := models.Event{
		GroupId:     gid,
		Title:       r.Form.Get("title"),
		Description: r.Form.Get("description"),
	}

	event.Date, err = time.Parse("2006-01-02", r.Form.Get("date"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	event.Picture = config.DEFAULT_AVATAR

	if len(r.MultipartForm.File["eventAvatar"]) != 0 {
		imageStorage := img.NewImageStorage(r, "eventAvatar")
		image, err := imageStorage.InitImage(config.AVATAR_SAVE_PATH)
		if err != nil && err != http.ErrMissingFile {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !image.CheckImgExtensionPermitted() {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ok, err := image.CheckImgSize(config.AVATAR_MAX_SIZE)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		image.Save()
		event.Picture = image.Name
	}

	eid, err := m.DB.InsertEvent(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.DB.InsertEventPicture(eid, event.Picture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ids, err := m.DB.GetGroupParticipants(gid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, id := range ids {
		err = m.DB.InsertEventParticipatingRequest(gid, eid, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	response.Data = fmt.Sprint(eid)
	response.Message = "Event successfully created!"
	js, err := json.Marshal(response)
	w.Write(js)
}
