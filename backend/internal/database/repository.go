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
	GetGroupFollowRequests(id int) ([]models.SocketMessage, error)
	GetEventRequests(id int) ([]models.SocketMessage, error)
	RemoveGroupFollowRequest(gid, sourceId int) error
	GetAllGroups() ([]models.Group, error)
	GetPost(id int) (models.Post, error)
	GetGroup(id int) (models.Group, error)
	GetGroupParticipants(id int) ([]int, error)
	GetGroupParticipantsByChat(id int) ([]int, error)
	CheckGroupIsPivate(id int) (bool, error)
	CheckChatExistence(fid, sid int) (bool, error)
	GetPostComments(id int) ([]models.Comment, error)
	GetAllPosts(userID, groupId int) ([]models.Post, error)
	GetAllGroupEvents(groupId int) ([]models.Event, error)
	GetAllChatMessages(chatId int) ([]models.ChatMessage, error)
	GetAllChats(uid int, group bool) ([]models.Chat, error)
	FollowUser(srcId, targetId int) error
	FollowGroup(uid, gid int) error
	GroupUnFollow(uid, gid int) error
	UnFollow(srcId, targetId int) error
	CheckPostAccessibility(userId int, post models.Post) (bool, error)
	InsertUserFollowRequest(srcId, targetId int) error
	InsertGroupFollowRequest(gid, creatorid, uid int, invite bool) error
	InsertPostShare(userId, postId int) error
	InsertPost(post models.Post) (int, error)
	InsertPostPicture(id int, path string) error
	InsertChat(groupChat bool) (int, error)
	InsertChatParticipant(chatId, uid int) error
	InsertChatMessage(message models.ChatMessage) error
	InsertGroup(group models.Group) (int, error)
	InsertGroupPicture(id int, path string) error
	InsertEvent(event models.Event) (int, error)
	InsertEventPicture(id int, path string) error
	InsertEventParticipater(eid, uid int, willAttend bool) error
	InsertEventParticipatingRequest(gid, eid, uid int) error
	GetEventParticipants(eventId int) ([]models.EventParticipant, error)
	RemoveFollowRequest(sourceId, destId int) error
	RemoveEventParticipatingRequest(eid, uid int) error
	GetUserFollowers(id int) ([]models.Follow, error)
	GetUserFollows(id int) ([]models.Follow, error)
	CheckFollowRequest(srcId, targetId int) (int, error)
	CheckGroupRequest(uid, gid int) (int, error)
	CheckProfileIsPivate(id int) (bool, error)
	CheckAlreadyUserFollowed(srcId, targetId int) (int, error)
	CheckAlreadyGroupFollowed(uid, gid int) (int, error)
	CheckAlreadyParticipating(eid, uid int) (int, error)
	CheckGroupInvite(uid, gid int) (int, error)
	CheckEmailExistence(email string) (int, error)
	CheckSessionExistence(token string) (int, error)
	UpdateSessionToken(token string, id int) error
	UpdateUserAvatar(id int, path string) error
	UpdateUserProfile(id int, aboutMe, nickname string) error
	UpdateUserPrivacy(id int, private bool) error
	GetUserHash(id int) (string, error)
}
