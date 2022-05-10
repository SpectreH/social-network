package image_storage

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

var permittedImgExtension []string = []string{"png", "jpg", "jpeg", "gif"}

// ImageStorage creates a custom Storage struct for images, embeds a r.Request object
type ImageStorage struct {
	r        *http.Request
	formName string
}

// Image is the model of saved image
type Image struct {
	Name   string
	Path   string
	Header *multipart.FileHeader
}

// NewStorage initializes a storage struct
func NewImageStorage(r *http.Request, formName string) *ImageStorage {
	return &ImageStorage{
		r,
		formName,
	}
}

// InitImage inits the image
func (s *ImageStorage) InitImage(savePath string) (Image, error) {
	image := Image{}
	var err error

	_, image.Header, err = s.r.FormFile(s.formName)
	if err != nil {
		return image, err
	}

	imageData := strings.Split(image.Header.Filename, ".")
	if len(imageData) < 2 {
		return image, errors.New("image doesn't have extension")
	}

	randBytes := make([]byte, 16)
	rand.Read(randBytes)

	image.Name = hex.EncodeToString(randBytes) + "." + imageData[len(imageData)-1]
	image.Path = savePath + image.Name

	return image, nil
}

// CheckImgExtensionPermitted checks if new img extension is permitted
func (i *Image) CheckImgExtensionPermitted() bool {
	imageData := strings.Split(i.Name, ".")
	for _, allowedExt := range permittedImgExtension {
		if imageData[len(imageData)-1] == allowedExt {
			return true
		}
	}

	return false
}

// Save saves image in given path
func (i *Image) Save() error {
	file, err := i.Header.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	out, err := os.Create(i.Path)
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(out, file)

	return nil
}

// CheckImgSize checks if img has permitted size
func (i *Image) CheckImgSize(size int64) (bool, error) {
	file, err := i.Header.Open()
	if err != nil {
		return false, err
	}
	defer file.Close()

	res, err := file.Seek(0, 2)
	if err != nil {
		return false, err
	}

	return res <= size, nil
}

// RemoveImage removes certain file from the directory
func (s *ImageStorage) RemoveImage(fileName string, path string) error {
	return os.Remove(path + fileName)
}
