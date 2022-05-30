package database

import "social-network/internal/models"

type DatabaseRepo interface {
	InsertUser(user models.User) (int, error)
	InesertSession(id int) error
	InsertProfileImage(id int, path string) error
	InsertPrivacySettings(id int) error
	GetUserData(id int) (models.User, error)
	GetUserAvatar(id int) (string, error)
	CheckEmailExistence(email string) (int, error)
	CheckSessionExistence(token string) (int, error)
	UpdateSessionToken(token string, id int) error
	UpdateUserAvatar(id int, path string) error
	UpdateUserProfile(id int, aboutMe, nickname string) error
	UpdateUserPrivacy(id int, private bool) error
	GetUserHash(id int) (string, error)
}
