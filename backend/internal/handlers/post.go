package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/internal/config"
	"social-network/internal/models"
	img "social-network/pkg/image-storage"
	"time"
)

func (m *Repository) GetPostContent(w http.ResponseWriter, r *http.Request) {
	_, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	pid, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	fmt.Println(pid)

	fullPost := models.PostInside{}

	if !response.OK {
		out, _ := json.MarshalIndent(response, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	fullPost.Post, err = m.DB.GetPost(pid)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "Post not founded!",
		}
	}

	if !response.OK {
		out, _ := json.MarshalIndent(response, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	fullPost.Comments, err = m.DB.GetPostComments(pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, _ := json.MarshalIndent(fullPost, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) PostNew(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var followers []models.Follow
	err = json.Unmarshal([]byte(r.Form.Get("followers")), &followers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post := models.Post{
		Content:   r.Form.Get("postContent"),
		ShareId:   r.Form.Get("postShare"),
		AuthId:    uid,
		CreatedAt: time.Now(),
	}

	if post.Content == "" {
		http.Error(w, "Content must be filled", http.StatusInternalServerError)
		return
	}

	if len(r.MultipartForm.File["postImage"]) != 0 {
		imageStorage := img.NewImageStorage(r, "postImage")
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
		post.Picture = image.Name
	}

	pid, err := m.DB.InsertPost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.DB.InsertPostPicture(pid, post.Picture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FormValidationResponse{OK: true, Message: "Post successfully created!"}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func (m *Repository) FollowerList(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	followers, err := m.DB.GetUserFollowers(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, _ := json.MarshalIndent(followers, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
