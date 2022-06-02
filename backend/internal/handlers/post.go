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

func (m *Repository) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	posts, err := m.DB.GetAllPosts(0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postsToShow := []models.PostInside{}

	for _, post := range posts {
		postToShow := models.PostInside{}

		var res bool
		res, err = m.DB.CheckPostAccessibility(uid, post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if res {
			postToShow.Post = post
			postToShow.Comments, err = m.DB.GetPostComments(post.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			postsToShow = append(postsToShow, postToShow)
		}
	}

	out, _ := json.MarshalIndent(postsToShow, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) GetPostContent(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	pid, err := getIdFromQuery(r, "id")
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

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
			Message: "Post not found!",
		}
	}

	res, err := m.DB.CheckPostAccessibility(uid, fullPost.Post)
	if !res {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "You don't have access to see this post!",
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

func (m *Repository) CommentNew(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	postId, err := strconv.Atoi(r.Form.Get("postId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment := models.Comment{
		AuthId:    uid,
		Content:   r.Form.Get("commentContent"),
		PostId:    postId,
		CreatedAt: time.Now(),
	}

	if comment.Content == "" {
		http.Error(w, "Comment must be filled", http.StatusInternalServerError)
		return
	}

	if len(r.MultipartForm.File["commentImage"]) != 0 {
		imageStorage := img.NewImageStorage(r, "commentImage")
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
		comment.Picture = image.Name
	}

	cid, err := m.DB.InsertComment(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.DB.InsertCommentPicture(cid, comment.Picture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FormValidationResponse{OK: true, Message: "Comment successfully created!"}
	js, _ := json.Marshal(response)
	w.Write(js)
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

	groupPost, err := strconv.ParseBool(r.Form.Get("groupPost"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var groupId int = 0
	if groupPost {
		groupId, err = strconv.Atoi(r.Form.Get("groupId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	shareId, err := strconv.Atoi(r.Form.Get("postShare"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post := models.Post{
		Content:   r.Form.Get("postContent"),
		ShareId:   shareId,
		AuthId:    uid,
		GroupId:   groupId,
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

	if post.ShareId == 2 {
		for _, f := range followers {
			err := m.DB.InsertPostShare(f.Id, pid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	err = m.DB.InsertPostPicture(pid, post.Picture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FormValidationResponse{OK: true, Message: "Post successfully created!", Data: fmt.Sprint("/post/", pid)}
	js, err := json.Marshal(response)
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
