package handlers

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"social-network/internal/models"
	img "social-network/pkg/image-storage"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
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

	imageStorage := img.NewImageStorage(r, "avatar")

	image, err := imageStorage.InitImage("./")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("OK"))
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
