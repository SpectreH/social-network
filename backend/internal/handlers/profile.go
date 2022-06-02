package handlers

import (
	"encoding/json"
	"net/http"
	"social-network/internal/models"
)

func (m *Repository) ProfileRequestToFollow(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	destId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	if destId == uid {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "You can't follow yourself!",
		}
	}

	if response.OK {
		res, err := m.DB.CheckFollowRequest(uid, destId)
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
		err = m.DB.InsertUserFollowRequest(uid, destId)
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

func (m *Repository) ProfileRemoveFollow(w http.ResponseWriter, r *http.Request) {
	uid, err := CheckSession(w, r)
	if err != nil {
		return
	}

	response := models.FormValidationResponse{
		OK: true,
	}

	sourceId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	if sourceId == uid {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "You can't remove yourself!",
		}
	}

	if response.OK {
		err = m.DB.UnFollow(sourceId, uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Message = "You successfully removed follower!"
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

	destId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	if destId == uid {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "You can't unfollow yourself!",
		}
	}

	if response.OK {
		err = m.DB.UnFollow(uid, destId)
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

	destId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	if destId == uid {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "You can't follow yourself!",
		}
	}

	if response.OK {
		private, err := m.DB.CheckProfileIsPivate(destId)
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
		res, err := m.DB.CheckAlreadyFollowed(uid, destId)
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
		err = m.DB.FollowUser(uid, destId)
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

	destId, err := getIdFromQuery(r)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: err.Error(),
		}
	}

	profile, err := m.DB.GetUserProfile(destId)
	if err != nil {
		response = models.FormValidationResponse{
			OK:      false,
			Message: "Profile with this id doesn't exist!",
		}
	}

	if !response.OK {
		out, _ := json.MarshalIndent(response, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	if destId == uid {
		profile.IsMyProfile = true
	}

	res, err := m.DB.CheckAlreadyFollowed(uid, destId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if res == 1 {
		profile.Following = true
	}

	if (profile.Private && profile.Following) || profile.IsMyProfile || !profile.Private {
		followers, err := m.DB.GetUserFollowers(destId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		follows, err := m.DB.GetUserFollows(destId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		posts, err := m.DB.GetAllPosts(destId)
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

		profile.Followers = append(followers, follows...)
		profile.Posts = postsToShow
	}

	out, _ := json.MarshalIndent(profile, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
