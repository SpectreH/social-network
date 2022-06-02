package database

import "social-network/internal/models"

type DatabaseRepo interface {
	InsertUser(user models.User) (int, error)
	InesertSession(id int) error
	InsertProfileImage(id int, path string) error
	InsertPrivacySettings(id int) error
	InsertComment(comment models.Comment) (int, error)
	InsertCommentPicture(id int, path string) error
	GetUserData(id int) (models.User, error)
	GetUserAvatar(id int) (string, error)
	GetUserProfile(id int) (models.UserProfile, error)
	GetUserFullName(id int) (string, error)
	GetUserFollowRequests(id int) ([]models.SocketMessage, error)
	GetPost(id int) (models.Post, error)
	GetPostComments(id int) ([]models.Comment, error)
	GetAllPosts(userID int) ([]models.Post, error)
	FollowUser(srcId, targetId int) error
	UnFollow(srcId, targetId int) error
	CheckPostAccessibility(userId int, post models.Post) (bool, error)
	InsertUserFollowRequest(srcId, targetId int) error
	InsertPostShare(userId, postId int) error
	InsertPost(post models.Post) (int, error)
	InsertPostPicture(id int, path string) error
	RemoveFollowRequest(sourceId, destId int) error
	GetUserFollowers(id int) ([]models.Follow, error)
	GetUserFollows(id int) ([]models.Follow, error)
	CheckFollowRequest(srcId, targetId int) (int, error)
	CheckProfileIsPivate(id int) (bool, error)
	CheckAlreadyFollowed(srcId, targetId int) (int, error)
	CheckEmailExistence(email string) (int, error)
	CheckSessionExistence(token string) (int, error)
	UpdateSessionToken(token string, id int) error
	UpdateUserAvatar(id int, path string) error
	UpdateUserProfile(id int, aboutMe, nickname string) error
	UpdateUserPrivacy(id int, private bool) error
	GetUserHash(id int) (string, error)
}
