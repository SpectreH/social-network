package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"social-network/internal/config"
	"social-network/internal/models"
	"social-network/pkg/forms"
	"strings"

	img "social-network/pkg/image-storage"

	"golang.org/x/crypto/bcrypt"
)

type ValidationResponse struct {
	OK      bool   `json:"ok"`
	Input   string `json:"input"`
	Message string `json:"message"`
}

func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := models.User{}
	for k, v := range r.Form {
		if len(v) == 0 {
			continue
		}

		value := v[0]
		if k == "password" {
			value = getHash([]byte(v[0]))
		}

		if err = setUserStructValue(&user, strings.Title(k), value); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	formValidator := forms.NewForm(r.Form)

	formValidator.Required("firstName", "lastName", "email", "password", "birthDate")
	formValidator.IsEmail("email")
	formValidator.MinLenght("password", 8)

	if r.MultipartForm.File["avatar"] != nil && formValidator.Valid() {
		imageStorage := img.NewImageStorage(r, "avatar")
		image, err := imageStorage.InitImage("./")
		if err != nil && err != http.ErrMissingFile {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !image.CheckImgExtensionPermitted() {
			formValidator.Errors.Add("avatar", "Only JPG, JPEG, PNG, GIF are allowed")
		}

		ok, err := image.CheckImgSize(config.AVATAR_MAX_SIZE)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !ok {
			formValidator.Errors.Add("avatar", "File size shoud be less than 5 MB")
		}

		if formValidator.Valid() {
			if err = image.Save(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	if !formValidator.Valid() {
		message, input := formValidator.Errors.GetFirst()

		response := ValidationResponse{
			OK:      false,
			Input:   input,
			Message: message,
		}

		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)
		return
	}

	response := ValidationResponse{OK: true}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func setUserStructValue(structure *models.User, field string, value string) error {
	r := reflect.ValueOf(structure)
	f := reflect.Indirect(r).FieldByName(field)

	if !f.IsValid() {
		return errors.New("non existing struct property")
	}

	if f.Kind() != reflect.Invalid {
		f.SetString(value)
	}

	return nil
}

// getHash generates hash from password
func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
