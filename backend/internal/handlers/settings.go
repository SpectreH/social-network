package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"social-network/internal/config"
	"social-network/internal/models"
	img "social-network/pkg/image-storage"
)

func (m *Repository) UpdatePrivacy(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var private bool
	if r.Form.Get("private") == "true" {
		private = true
	}

	err = m.DB.UpdateUserPrivacy(uid, private)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FormValidationResponse{
		OK:      true,
		Message: "Profile privacy successfully updated!",
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func (m *Repository) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aboutMe := r.Form.Get("aboutMe")
	nickname := r.Form.Get("nickname")
	m.DB.UpdateUserProfile(uid, aboutMe, nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FormValidationResponse{
		OK:      true,
		Message: "Profile settings successfully updated!",
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func (m *Repository) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	imageStorage := img.NewImageStorage(r, "avatar")
	image, err := imageStorage.InitImage(config.AVATAR_SAVE_PATH)
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FormValidationResponse{
		OK:      true,
		Message: "Success",
	}

	if !image.CheckImgExtensionPermitted() {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "Only JPG, JPEG, PNG, GIF are allowed",
		}
	}

	ok, err := image.CheckImgSize(config.AVATAR_MAX_SIZE)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "File size shoud be less than 5 MB",
		}
	}

	path, err := m.DB.GetUserAvatar(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if path != config.DEFAULT_AVATAR {
		err = os.Remove(fmt.Sprintf("./images/%s", path))
	}

	err = m.DB.UpdateUserAvatar(uid, image.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = image.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response = models.FormValidationResponse{
		OK:   true,
		Data: config.AVATAR_PATH_URL + image.Name,
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
