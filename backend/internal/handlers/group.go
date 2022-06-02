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

func (m *Repository) GroupRequestToFollow(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	groupId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	if response.OK {
		res, err := m.DB.CheckAlreadyGroupFollowed(uid, groupId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if res != 0 {
			response = models.FormValidationResponse{
				OK:      false,
				Message: "You already following that group!",
			}
		}
	}

	if response.OK {
		res, err := m.DB.CheckGroupRequest(uid, groupId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		group, err := m.DB.GetGroup(groupId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = m.DB.InsertGroupFollowRequest(groupId, group.CreatorId, uid, false)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Message = "Follow request successfully sended!"
	}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}

func (m *Repository) GroupUnFollow(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	groupId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	if response.OK {
		err = m.DB.GroupUnFollow(uid, groupId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Message = "You stopped following this group!"
	}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}

func (m *Repository) GroupFollow(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	groupId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	if response.OK {
		private, err := m.DB.CheckGroupIsPivate(groupId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if private {
			response = models.FormValidationResponse{
				OK:      false,
				Message: "You must request to follow this group!",
			}
		}
	}

	if response.OK {
		res, err := m.DB.CheckAlreadyGroupFollowed(uid, groupId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if res != 0 {
			response = models.FormValidationResponse{
				OK:      false,
				Message: "You already following that group!",
			}
		}
	}

	if response.OK {
		err = m.DB.FollowGroup(uid, groupId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Message = "You now following this group!"
	}

	out, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}

func (m *Repository) GetGroup(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	groupId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	group, err := m.DB.GetGroup(groupId)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "Group with this id doesn't exist!",
		}
	}

	if !response.OK {
		out, _ := json.MarshalIndent(response, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	if group.CreatorId == uid {
		group.IsMyGroup = true
		group.Following = true
	} else {
		res, err := m.DB.CheckAlreadyGroupFollowed(uid, groupId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if res == 1 {
			group.Following = true
		}
	}

	posts, err := m.DB.GetAllPosts(0, groupId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, post := range posts {
		postToShow := models.PostInside{}
		postToShow.Post = post
		postToShow.Comments, err = m.DB.GetPostComments(post.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		group.Posts = append(group.Posts, postToShow)
	}

	out, _ := json.MarshalIndent(group, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	_, err := CheckSession(w, r)
	if err != nil {
		return
	}

	groups, err := m.DB.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, _ := json.MarshalIndent(groups, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) GroupNew(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shareId, err := strconv.Atoi(r.Form.Get("currentShareSettings"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var isPrivate bool
	if shareId == 1 {
		isPrivate = true
	}

	group := models.Group{
		Title:       r.Form.Get("title"),
		Description: r.Form.Get("description"),
		Picture:     config.DEFAULT_AVATAR,
		Private:     isPrivate,
		CreatorId:   uid,
		CreatedAt:   time.Now(),
	}

	if len(r.MultipartForm.File["groupAvatar"]) != 0 {
		imageStorage := img.NewImageStorage(r, "groupAvatar")
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
		group.Picture = image.Name
	}

	cid, err := m.DB.InsertChat(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	group.ChatId = cid

	gid, err := m.DB.InsertGroup(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.DB.InsertGroupPicture(gid, group.Picture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FormValidationResponse{OK: true, Message: "Group successfully created!", Data: fmt.Sprint("/group/", gid)}
	js, err := json.Marshal(response)
	w.Write(js)
}
