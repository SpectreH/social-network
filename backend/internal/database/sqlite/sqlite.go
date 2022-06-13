package sqlite

import (
	"database/sql"
	"fmt"
	"social-network/internal/config"
	"social-network/internal/models"
	"strings"
	"time"
)

// InsertUser inserts a new user into database
func (m *sqliteDBRepo) InsertUser(user models.User) (int, error) {
	var userID int

	query := `insert into users (first_name, last_name, email, birth_date, nickname, about_me, password) values ($1, $2, $3, $4, $5, $6, $7) returning id;`
	row := m.DB.QueryRow(query, user.FirstName, user.LastName, user.Email, user.BirthDate, user.Nickname, user.AboutMe, user.Password)
	err := row.Scan(&userID)

	if err != nil {
		return -1, err
	}

	return userID, nil
}

// InesertSession inserts a new user session cell into database
func (m *sqliteDBRepo) InesertSession(id int) error {
	query := `insert into user_sessions (user_id) values ($1);`
	_, err := m.DB.Exec(query, id)

	return err
}

// InsertProfileImage inserts a new user profile image cell into database
func (m *sqliteDBRepo) InsertProfileImage(id int, path string) error {
	query := `insert into user_profile_images (user_id, path) values ($1, $2);`
	_, err := m.DB.Exec(query, id, path)

	return err
}

// InsertPrivacySettings inserts a new user privacy settings
func (m *sqliteDBRepo) InsertPrivacySettings(id int) error {
	query := `insert into user_privacy_settings (user_id, private_account) values ($1, $2);`
	_, err := m.DB.Exec(query, id, false)

	return err
}

// InsertComment inserts new comment
func (m *sqliteDBRepo) InsertComment(comment models.Comment) (int, error) {
	var id int

	query := `insert into comments (user_id, post_id, content, created_at) values ($1, $2, $3, $4) returning id;`
	err := m.DB.QueryRow(query, comment.AuthId, comment.PostId, comment.Content, comment.CreatedAt).Scan(&id)

	return id, err
}

// InsertCommentPicture inserts comment picture
func (m *sqliteDBRepo) InsertCommentPicture(id int, path string) error {
	query := `insert into comment_images (comment_id, path) values ($1, $2);`
	_, err := m.DB.Exec(query, id, path)

	return err
}

// InsertChatParticipant inserts new chat participant
func (m *sqliteDBRepo) InsertChatParticipant(chatId, uid int) error {
	query := `insert into chat_participants (chat_id, participant_id) values ($1, $2);`
	_, err := m.DB.Exec(query, chatId, uid)

	return err
}

// InsertChatMessage inserts new chat message
func (m *sqliteDBRepo) InsertChatMessage(message models.ChatMessage) error {
	query := `insert into chat_messages (chat_id, author_id, content, created_at) values ($1, $2, $3, $4);`
	_, err := m.DB.Exec(query, message.ChatId, message.AuthorId, message.Content, time.Now())

	return err
}

// GetUserFullName gets users's full name
func (m *sqliteDBRepo) GetUserFullName(id int) (string, error) {
	var fn, ln string

	query := `select u.first_name, u.last_name from users u where u.id = $1`
	err := m.DB.QueryRow(query, id).Scan(&fn, &ln)

	return fmt.Sprintf("%s %s", fn, ln), err
}

// GetUserProfile gets user's profile
func (m *sqliteDBRepo) GetUserProfile(id int) (models.UserProfile, error) {
	var profile models.UserProfile

	query := `select u.id, u.first_name, u.last_name, u.email, u.birth_date, u.nickname, u.about_me, upi.path, ups.private_account,
  (SELECT COUNT(*) FROM followers WHERE user_id = u.id),
  (SELECT COUNT(*) FROM followers WHERE follower_id = u.id),
  (SELECT COUNT(*) FROM posts WHERE user_id = u.id)
  from users u 
	JOIN user_profile_images upi ON upi.user_id = u.id 
	JOIN user_privacy_settings ups ON ups.id = u.id     
	where u.id = $1`

	err := m.DB.QueryRow(query, id).Scan(&profile.Id, &profile.FirstName, &profile.LastName,
		&profile.Email, &profile.BirthDate, &profile.Nickname,
		&profile.AboutMe, &profile.Avatar, &profile.Private,
		&profile.TotalFollowers, &profile.TotalFollows, &profile.TotalPosts)

	profile.Avatar = config.AVATAR_PATH_URL + profile.Avatar

	return profile, err
}

// RemoveGroupFollowRequest removes a follow request of certain user at certain group
func (m *sqliteDBRepo) RemoveGroupFollowRequest(gid, sourceId int) error {
	query := `DELETE FROM group_follow_requests WHERE group_id = $1 AND user_id = $2`
	_, err := m.DB.Exec(query, gid, sourceId)

	return err
}

// InsertGroupFollowRequest inserts group's follow request with status pending
func (m *sqliteDBRepo) InsertGroupFollowRequest(gid, creatorid, uid int, invite bool) error {
	query := `insert into group_follow_requests (request_status_id, group_id, creator_id, user_id, invite, requested_at) values ($1, $2, $3, $4, $5, $6);`
	_, err := m.DB.Exec(query, 1, gid, creatorid, uid, invite, time.Now())

	return err
}

// InsertUserFollowRequest inserts uses's follow request with status pending
func (m *sqliteDBRepo) InsertUserFollowRequest(srcId, targetId int) error {
	query := `insert into follow_requests (request_status_id, follow_from, follow_to, requested_at) values ($1, $2, $3, $4);`
	_, err := m.DB.Exec(query, 1, srcId, targetId, time.Now())

	return err
}

// InserPostShare inserts a record of post access for certain user
func (m *sqliteDBRepo) InsertPostShare(userId, postId int) error {
	query := `insert into post_shares (post_id, user_id) values ($1, $2);`
	_, err := m.DB.Exec(query, postId, userId)

	return err
}

// InsertPost inserts new post
func (m *sqliteDBRepo) InsertPost(post models.Post) (int, error) {
	var id int

	query := `insert into posts (group_id, user_id, share_id, content, created_at) values ($1, $2, $3, $4, $5) returning id;`
	err := m.DB.QueryRow(query, post.GroupId, post.AuthId, post.ShareId, post.Content, post.CreatedAt).Scan(&id)

	return id, err
}

// InsertPostPicture inserts post picture
func (m *sqliteDBRepo) InsertPostPicture(id int, path string) error {
	query := `insert into post_images (post_id, path) values ($1, $2);`
	_, err := m.DB.Exec(query, id, path)

	return err
}

// InsertChat inserts new chat
func (m *sqliteDBRepo) InsertChat(groupChat bool) (int, error) {
	var id int

	query := `insert into chats (group_chat) values ($1) returning id;`
	err := m.DB.QueryRow(query, groupChat).Scan(&id)

	return id, err
}

// InsertGroup inserts new group
func (m *sqliteDBRepo) InsertGroup(group models.Group) (int, error) {
	var id int

	query := `insert into groups (chat_id, creator_id, title, description, private, created_at) values ($1, $2, $3, $4, $5, $6) returning id;`
	err := m.DB.QueryRow(query, group.ChatId, group.CreatorId, group.Title, group.Description, group.Private, group.CreatedAt).Scan(&id)

	return id, err
}

// InsertGroupPicture inserts group picture
func (m *sqliteDBRepo) InsertGroupPicture(id int, path string) error {
	query := `insert into group_images (group_id, path) values ($1, $2);`
	_, err := m.DB.Exec(query, id, path)

	return err
}

// InsertEvent inserts new event
func (m *sqliteDBRepo) InsertEvent(event models.Event) (int, error) {
	var id int

	query := `insert into group_events (group_id, title, description, date, created_at) values ($1, $2, $3, $4, $5) returning id;`
	err := m.DB.QueryRow(query, event.GroupId, event.Title, event.Description, event.Date, time.Now()).Scan(&id)

	return id, err
}

// InsertEventPicture inserts event picture
func (m *sqliteDBRepo) InsertEventPicture(id int, path string) error {
	query := `insert into group_event_images (event_id, path) values ($1, $2);`
	_, err := m.DB.Exec(query, id, path)

	return err
}

// InsertEventParticipatingRequest inserts group's participating
func (m *sqliteDBRepo) InsertEventParticipatingRequest(gid, eid, uid int) error {
	query := `insert into group_event_request (request_status_id, group_id, event_id, user_id, requested_at) values ($1, $2, $3, $4, $5);`
	_, err := m.DB.Exec(query, 1, gid, eid, uid, time.Now())

	return err
}

// InsertEventParticipater inserts event participant
func (m *sqliteDBRepo) InsertEventParticipater(eid, uid int, willAttend bool) error {
	query := `insert into group_event_participants (participant_id, event_id, will_attend) values ($1, $2, $3);`
	_, err := m.DB.Exec(query, uid, eid, willAttend)

	return err
}

// CheckChatExistence checks if chat exists between two users
func (m *sqliteDBRepo) CheckChatExistence(fid, sid int) (bool, error) {
	var res int

	query := `
		SELECT COUNT(*)
		FROM chats c 
		INNER JOIN chat_participants p1 ON p1.chat_id = c.id AND p1.participant_id = $1
		INNER JOIN chat_participants p2 ON p2.chat_id = c.id AND p2.participant_id = $2
	`
	err := m.DB.QueryRow(query, fid, sid).Scan(&res)

	return res == 1, err
}

// CheckPostAccessibility checks if user has access to see certain post
func (m *sqliteDBRepo) CheckPostAccessibility(userId int, post models.Post) (bool, error) {
	var res int

	if userId == post.AuthId {
		return true, nil
	}

	switch post.ShareId {
	case 1:
		query := `select COUNT(*) from followers WHERE follower_id = $1 AND user_id= $2;`
		err := m.DB.QueryRow(query, userId, post.AuthId).Scan(&res)
		return res != 0, err
	case 2:
		query := `select COUNT(*) from post_shares WHERE post_id = $1 AND user_id= $2;`
		err := m.DB.QueryRow(query, post.Id, userId).Scan(&res)
		return res != 0, err
	case 3:
		query := `select COUNT(*) from group_membership WHERE group_id = $1 AND user_id= $2;`
		err := m.DB.QueryRow(query, post.GroupId, userId).Scan(&res)
		return res != 0, err
	default:
		return true, nil
	}
}

// CheckGroupRequest checks if follow request already exists
func (m *sqliteDBRepo) CheckGroupRequest(uid, gid int) (int, error) {
	var res int
	query := `select COUNT(*) from group_follow_requests WHERE user_id = $1 AND group_id = $2;`
	err := m.DB.QueryRow(query, uid, gid).Scan(&res)

	return res, err
}

// CheckFollowRequest checks if follow request already exists
func (m *sqliteDBRepo) CheckFollowRequest(srcId, targetId int) (int, error) {
	var res int
	query := `select COUNT(*) from follow_requests WHERE follow_from = $1 AND follow_to= $2;`
	err := m.DB.QueryRow(query, srcId, targetId).Scan(&res)

	return res, err
}

// CheckGroupIsPivate checks if group has private group
func (m *sqliteDBRepo) CheckGroupIsPivate(id int) (bool, error) {
	var res bool

	query := `select g.private from groups g where g.id = $1;`
	err := m.DB.QueryRow(query, id).Scan(&res)

	return res, err
}

// CheckProfileIsPivate checks if user has private account
func (m *sqliteDBRepo) CheckProfileIsPivate(id int) (bool, error) {
	var res bool

	query := `select ups.private_account from user_privacy_settings ups where ups.user_id = $1;`
	err := m.DB.QueryRow(query, id).Scan(&res)

	return res, err
}

// CheckGroupInvite checks if user has invite
func (m *sqliteDBRepo) CheckGroupInvite(uid, gid int) (int, error) {
	var res int
	query := `select COUNT(*) from group_follow_requests WHERE user_id = $1 AND group_id = $2 AND invite = true;`
	err := m.DB.QueryRow(query, uid, gid).Scan(&res)

	return res, err
}

// CheckAlreadyParticipating checks if user already participating certain event
func (m *sqliteDBRepo) CheckAlreadyParticipating(eid, uid int) (int, error) {
	var res int
	query := `select COUNT(*) from group_event_participants WHERE participant_id = $1 AND event_id = $2;`
	err := m.DB.QueryRow(query, uid, eid).Scan(&res)

	return res, err
}

// CheckAlreadyGroupFollowed checks if user already followed certain group
func (m *sqliteDBRepo) CheckAlreadyGroupFollowed(uid, gid int) (int, error) {
	var res int
	query := `select COUNT(*) from group_membership WHERE user_id = $1 AND group_id = $2;`
	err := m.DB.QueryRow(query, uid, gid).Scan(&res)

	return res, err
}

// CheckAlreadyUserFollowed checks if user already followed certain user
func (m *sqliteDBRepo) CheckAlreadyUserFollowed(srcId, targetId int) (int, error) {
	var res int
	query := `select COUNT(*) from followers WHERE follower_id = $1 AND user_id = $2;`
	err := m.DB.QueryRow(query, srcId, targetId).Scan(&res)

	return res, err
}

// FollowGroup makes a record with follow
func (m *sqliteDBRepo) FollowGroup(uid, gid int) error {
	query := `insert into group_membership (user_id, group_id, joined_at) values ($1, $2, $3);`
	_, err := m.DB.Exec(query, uid, gid, time.Now())

	return err
}

// GroupUnFollow deletes record about following
func (m *sqliteDBRepo) GroupUnFollow(uid, gid int) error {
	query := `delete from group_membership where user_id = $1 AND group_id = $2;`
	_, err := m.DB.Exec(query, uid, gid)

	return err
}

// FollowUser makes a record with follow
func (m *sqliteDBRepo) FollowUser(srcId, targetId int) error {
	query := `insert into followers (user_id, follower_id, followed_at) values ($1, $2, $3);`
	_, err := m.DB.Exec(query, targetId, srcId, time.Now())

	res, err := m.CheckChatExistence(srcId, targetId)
	if err != nil {
		return err
	}

	if !res {
		cid, err := m.InsertChat(false)
		if err != nil {
			return err
		}

		err = m.InsertChatParticipant(cid, srcId)
		if err != nil {
			return err
		}

		err = m.InsertChatParticipant(cid, targetId)
		if err != nil {
			return err
		}
	}

	return err
}

// UnFollow deletes record about following
func (m *sqliteDBRepo) UnFollow(srcId, targetId int) error {
	query := `delete from followers where follower_id = $1 AND user_id = $2;`
	_, err := m.DB.Exec(query, srcId, targetId)

	return err
}

// GetPost gets all data about post
func (m *sqliteDBRepo) GetPost(id int) (models.Post, error) {
	var res models.Post

	query := `SELECT p.id, p.user_id, p.group_id, IFNULL(g.title, ''), IFNULL(gi.path, ''), u.first_name, u.last_name, ufi.path, p.share_id, p.content, pi.path, p.created_at FROM posts p
	JOIn post_images pi ON pi.post_id = p.id
	JOIn users u ON u.id = p.user_id
	JOIn user_profile_images ufi ON ufi.user_id = p.user_id
	Left OUTER JOIN groups g ON g.id = p.group_id
	Left OUTER JOIn group_images gi ON gi.group_id = p.group_id
	WHEre p.id = $1;`

	err := m.DB.QueryRow(query, id).Scan(&res.Id, &res.AuthId, &res.GroupId, &res.GroupTitle, &res.GroupAvatar, &res.FirstName, &res.LastName, &res.Avatar, &res.ShareId, &res.Content, &res.Picture, &res.CreatedAt)

	if res.GroupId != 0 {
		res.GroupAvatar = config.AVATAR_PATH_URL + res.GroupAvatar
	}

	if res.Picture != "" {
		res.Picture = config.AVATAR_PATH_URL + res.Picture
	}

	res.Avatar = config.AVATAR_PATH_URL + res.Avatar

	return res, err
}

// GetGroupParticipants gets all group followers id
func (m *sqliteDBRepo) GetGroupParticipants(id int) ([]int, error) {
	var participants []int

	query := `SELECT user_id FROM group_membership WHERE group_id = $1;`

	rows, err := m.DB.Query(query, id)
	if err != nil && err != sql.ErrNoRows {
		return participants, err
	}

	for rows.Next() {
		var participant int

		if rows.Scan(&participant) != nil {
			return participants, err
		}

		participants = append(participants, participant)
	}

	return participants, nil
}

// GetGroup gets all data about group
func (m *sqliteDBRepo) GetGroup(id int) (models.Group, error) {
	var group models.Group

	query := `SELECT g.id, g.chat_id, g.creator_id, g.title, g.description, g.private, gi.path, 
	(SELECT COUNT(*) FROM posts p WHERE p.group_id = g.id), 
	(SELECT COUNT(*) FROM group_membership gm WHERE gm.group_id = g.id), g.created_at 
	FROM groups g
	JOIN group_images gi ON gi.group_id = g.id
	WHERE g.id = $1;`

	err := m.DB.QueryRow(query, id).Scan(&group.Id, &group.ChatId, &group.CreatorId, &group.Title, &group.Description, &group.Private, &group.Picture, &group.TotalPosts, &group.TotalFollowers, &group.CreatedAt)

	group.Picture = config.AVATAR_PATH_URL + group.Picture

	return group, err
}

// GetEventParticipants gets all event participants from database
func (m *sqliteDBRepo) GetEventParticipants(eventId int) ([]models.EventParticipant, error) {
	var participants []models.EventParticipant

	query := `SELECT gep.id, gep.participant_id, gep.event_id, gep.will_attend FROM group_event_participants gep WHERE gep.event_id = $1`

	rows, err := m.DB.Query(query, eventId)
	if err != nil && err != sql.ErrNoRows {
		return participants, err
	}

	for rows.Next() {
		var participant models.EventParticipant

		if rows.Scan(&participant.Id, &participant.ParticipantId, &participant.EventId, &participant.WillAttend) != nil {
			return participants, err
		}

		participants = append(participants, participant)
	}

	return participants, err
}

// GetAllChats gets all group/direct chats
func (m *sqliteDBRepo) GetAllChats(uid int, group bool) ([]models.Chat, error) {
	var chats []models.Chat
	var query string

	if group {
		query = `SELECt g.chat_id, gm.group_id, g.title, gi.path FROM group_membership gm
	JOIN groups g ON g.id = gm.group_id
	JOIN group_images gi ON gi.group_id = gm.group_id
	WHERE gm.user_id = $1;`
	} else {
		query = `SELECT c.id, u.id, u.first_name || ' ' || u.last_name, upi.path
 	FROM chats c 
	INNER JOIN chat_participants p1 ON p1.chat_id = c.id AND p1.participant_id = $1
	INNER JOIN chat_participants p2 ON p2.chat_id = c.id AND p2.participant_id != $1 
	JOIN users u ON p2.participant_id = u.id
	JOIN user_profile_images upi ON upi.user_id = p2.participant_id`
	}

	rows, err := m.DB.Query(query, uid)
	if err != nil && err != sql.ErrNoRows {
		return chats, err
	}

	for rows.Next() {
		var chat models.Chat
		if err := rows.Scan(&chat.Id, &chat.DestId, &chat.Name, &chat.Avatar); err != nil {
			return chats, err
		}

		chat.Avatar = config.AVATAR_PATH_URL + chat.Avatar

		chat.Messages, err = m.GetAllChatMessages(chat.Id)
		if err != nil {
			return chats, err
		}

		chat.GroupChat = group

		chats = append(chats, chat)
	}

	return chats, nil
}

// GetAllChatMessages gets all messages of certain chat
func (m *sqliteDBRepo) GetAllChatMessages(chatId int) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage

	query := `SELECT cm.id, cm.chat_id, cm.author_id, cm.content, u.first_name, upi.path, cm.created_at FROM chat_messages cm 
	JOIN user_profile_images upi on upi.user_id = cm.author_id
	JOIN users u on u.id = cm.author_id
	WHERE cm.chat_id = $1`

	rows, err := m.DB.Query(query, chatId)
	if err != nil && err != sql.ErrNoRows {
		return messages, err
	}

	for rows.Next() {
		var message models.ChatMessage

		if rows.Scan(&message.Id, &message.ChatId, &message.AuthorId, &message.Content, &message.AuthorName, &message.ProfilePicture, &message.CreatedAt) != nil {
			return messages, err
		}

		message.ProfilePicture = config.AVATAR_PATH_URL + message.ProfilePicture

		messages = append(messages, message)
	}

	return messages, err
}

// GetAllGroupEvents gets all group events from database
func (m *sqliteDBRepo) GetAllGroupEvents(groupId int) ([]models.Event, error) {
	var events []models.Event

	query := `SELECT ge.id, ge.group_id, ge.title, ge.description, ge.date, gei.path FROM group_events ge
	JOIN group_event_images gei ON gei.event_id = ge.id
	WHERE ge.group_id = $1`

	rows, err := m.DB.Query(query, groupId)
	if err != nil && err != sql.ErrNoRows {
		return events, err
	}

	for rows.Next() {
		var event models.Event

		if rows.Scan(&event.Id, &event.GroupId, &event.Title, &event.Description, &event.Date, &event.Picture) != nil {
			return events, err
		}

		event.Picture = config.AVATAR_PATH_URL + event.Picture

		event.Participants, err = m.GetEventParticipants(event.Id)
		if err != nil {
			return events, err
		}

		events = append(events, event)
	}

	return events, err
}

// GetAllGroups gets all groups from database
func (m *sqliteDBRepo) GetAllGroups() ([]models.Group, error) {
	var groups []models.Group

	query := `SELECT g.id, g.chat_id, g.creator_id, g.title, g.description, g.private, gi.path, (SELECT COUNT(*) FROM posts p WHERE p.group_id = g.id), (SELECT COUNT(*) FROM group_membership gm WHERE gm.group_id = g.id), g.created_at FROM groups g
	JOIN group_images gi ON gi.group_id = g.id`

	rows, err := m.DB.Query(query)
	if err != nil && err != sql.ErrNoRows {
		return groups, err
	}

	for rows.Next() {
		var group models.Group

		if rows.Scan(&group.Id, &group.ChatId, &group.CreatorId, &group.Title, &group.Description, &group.Private, &group.Picture, &group.TotalPosts, &group.TotalFollowers, &group.CreatedAt) != nil {
			return groups, err
		}

		group.Picture = config.AVATAR_PATH_URL + group.Picture

		groups = append(groups, group)
	}

	return groups, err
}

// GetAllPosts gets all posts from database (option: of certain user or group)
func (m *sqliteDBRepo) GetAllPosts(userID, groupId int) ([]models.Post, error) {
	var posts []models.Post

	var rows *sql.Rows
	var err error

	if userID != 0 {
		query := `SELECT p.id, p.user_id, p.group_id, IFNULL(g.title, ''), IFNULL(gi.path, ''), u.first_name, u.last_name, ufi.path, p.share_id, p.content, pi.path, p.created_at FROM posts p
		JOIn post_images pi ON pi.post_id = p.id
		JOIn users u ON u.id = p.user_id
		JOIn user_profile_images ufi ON ufi.user_id = p.user_id
		Left OUTER JOIN groups g ON g.id = p.group_id
		Left OUTER JOIn group_images gi ON gi.group_id = p.group_id
		WHERE p.user_id = $1
		ORDER BY p.created_at DESC;`
		rows, err = m.DB.Query(query, userID)
	} else if groupId != 0 {
		query := `SELECT p.id, p.user_id, p.group_id, IFNULL(g.title, ''), IFNULL(gi.path, ''), u.first_name, u.last_name, ufi.path, p.share_id, p.content, pi.path, p.created_at FROM posts p
		JOIn post_images pi ON pi.post_id = p.id
		JOIn users u ON u.id = p.user_id
		JOIn user_profile_images ufi ON ufi.user_id = p.user_id
		Left OUTER JOIN groups g ON g.id = p.group_id
		Left OUTER JOIn group_images gi ON gi.group_id = p.group_id
		WHERE p.group_id = $1
		ORDER BY p.created_at DESC;`
		rows, err = m.DB.Query(query, groupId)
	} else {
		query := `SELECT p.id, p.user_id, p.group_id, IFNULL(g.title, ''), IFNULL(gi.path, ''), u.first_name, u.last_name, ufi.path, p.share_id, p.content, pi.path, p.created_at FROM posts p
		JOIn post_images pi ON pi.post_id = p.id
		JOIn users u ON u.id = p.user_id
		JOIn user_profile_images ufi ON ufi.user_id = p.user_id
		Left OUTER JOIN groups g ON g.id = p.group_id
		Left OUTER JOIn group_images gi ON gi.group_id = p.group_id
		ORDER BY p.created_at DESC;`
		rows, err = m.DB.Query(query)
	}

	if err != nil && err != sql.ErrNoRows {
		return posts, err
	}

	for rows.Next() {
		var post models.Post

		if rows.Scan(&post.Id, &post.AuthId, &post.GroupId, &post.GroupTitle, &post.GroupAvatar, &post.FirstName, &post.LastName, &post.Avatar, &post.ShareId, &post.Content, &post.Picture, &post.CreatedAt) != nil {
			return posts, err
		}

		if post.GroupId != 0 {
			post.GroupAvatar = config.AVATAR_PATH_URL + post.GroupAvatar
		}

		if post.Picture != "" {
			post.Picture = config.AVATAR_PATH_URL + post.Picture
		}

		post.Avatar = config.AVATAR_PATH_URL + post.Avatar

		posts = append(posts, post)
	}

	return posts, nil
}

// GetPostComments gets all post comments
func (m *sqliteDBRepo) GetPostComments(id int) ([]models.Comment, error) {
	var comments []models.Comment

	query := `SELECT c.id, c.user_id, u.first_name, u.last_name, ufi.path, c.content, c.created_at, ci.path FROM comments c
	JOIN comment_images ci ON ci.comment_id= c.id
	JOIN users u ON u.id = c.user_id  
	JOIn user_profile_images ufi ON ufi.user_id = c.user_id
	WHEre c.post_id = $1
	ORDER BY c.created_at DESC;`

	rows, err := m.DB.Query(query, id)
	if err != nil && err != sql.ErrNoRows {
		return comments, err
	}

	for rows.Next() {
		var comment models.Comment

		if rows.Scan(&comment.Id, &comment.AuthId, &comment.FirstName, &comment.LastName, &comment.Avatar, &comment.Content, &comment.CreatedAt, &comment.Picture) != nil {
			return comments, err
		}

		if comment.Picture != "" {
			comment.Picture = config.AVATAR_PATH_URL + comment.Picture
		}

		comment.Avatar = config.AVATAR_PATH_URL + comment.Avatar

		comments = append(comments, comment)
	}

	return comments, nil
}

// GetGroupFollowRequests gets all event requests of certain user
func (m *sqliteDBRepo) GetEventRequests(id int) ([]models.SocketMessage, error) {
	var requests []models.SocketMessage

	query := `SELECT ger.group_id, g.title, gi.path, ge.date, ge.title, ger.event_id, ger.user_id, ger.requested_at FROM group_event_request ger
	JOIN groups g ON g.id = ger.group_id
	JOIN group_images gi ON gi.group_id = ger.group_id
	JOIN group_events ge ON ge.id = ger.event_id
	WHERE ger.user_id = $1`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return requests, err
	}

	for rows.Next() {
		var request models.SocketMessage

		if rows.Scan(&request.GroupId, &request.SourceName, &request.Avatar, &request.EventDate, &request.EventName, &request.EventId, &request.To, &request.Date) != nil {
			return requests, err
		}

		request.EventDate = strings.Split(request.EventDate, "T")[0]

		request.Message = fmt.Sprintf("Has new event %s what will be held on %s", request.EventName, request.EventDate)
		request.Avatar = config.AVATAR_PATH_URL + request.Avatar
		request.Type = config.NEW_EVENT_REQUEST_TYPE

		requests = append(requests, request)
	}

	return requests, nil
}

// GetUserFollowRequests gets all group follow requests of certain user
func (m *sqliteDBRepo) GetGroupFollowRequests(id int) ([]models.SocketMessage, error) {
	var requests []models.SocketMessage

	query := `SELECT gfr.user_id, u.first_name, u.last_name, upi.path, gfr.group_id, gfr.creator_id, ugc.first_name, ugc.last_name, g.title, gfr.requested_at, gfr.invite FROM group_follow_requests gfr
	JOIN groups g ON g.id = gfr.group_id
	JOIN users u ON u.id = gfr.user_id
	JOIN users ugc ON ugc.id = gfr.creator_id
	JOIN user_profile_images upi ON upi.user_id = gfr.user_id
	WHERE (gfr.creator_id = $1 AND gfr.invite = false) OR (gfr.user_id = $1 AND gfr.invite = true);`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return requests, err
	}

	for rows.Next() {
		var request models.SocketMessage
		var invite bool
		var fn, ln, cfn, cln string

		if rows.Scan(&request.Source, &fn, &ln, &request.Avatar, &request.GroupId, &request.To, &cfn, &cln, &request.GroupName, &request.Date, &invite) != nil {
			return requests, err
		}

		request.SourceName = fmt.Sprintf("%s %s", fn, ln)
		request.Message = config.GROUP_FOLLOW_REQUEST_MESSAGE + request.GroupName

		if invite {
			request.SourceName = fmt.Sprintf("%s %s", cfn, cln)
			request.Message = config.GROUP_INVITE_MESSAGE + request.GroupName
		}

		request.Avatar = config.AVATAR_PATH_URL + request.Avatar
		request.Type = config.GROUP_FOLLOW_REQUEST_TYPE

		requests = append(requests, request)
	}

	return requests, nil
}

// GetUserFollowRequests gets all follow requests of certain user
func (m *sqliteDBRepo) GetUserFollowRequests(id int) ([]models.SocketMessage, error) {
	var requests []models.SocketMessage

	query := `SELECT fr.follow_from, u.first_name, u.last_name, upi.path, fr.follow_to, fr.requested_at FROM follow_requests fr
	JOIN users u ON u.id = fr.follow_from
	JOIN user_profile_images upi ON upi.user_id = fr.follow_from
	WHERE fr.follow_to = $1;`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return requests, err
	}

	for rows.Next() {
		var request models.SocketMessage
		var fn, ln string

		if rows.Scan(&request.Source, &fn, &ln, &request.Avatar, &request.To, &request.Date) != nil {
			return requests, err
		}

		request.SourceName = fmt.Sprintf("%s %s", fn, ln)
		request.Message = config.FOLLOW_REQUEST_MESSAGE
		request.Avatar = config.AVATAR_PATH_URL + request.Avatar
		request.Type = config.FOLLOW_REQUEST_TYPE

		requests = append(requests, request)
	}

	return requests, nil
}

// GetUserData gets information about user
func (m *sqliteDBRepo) GetUserData(id int) (models.User, error) {
	var user models.User

	query := `select u.id, u.first_name, u.last_name, u.email, u.birth_date, u.nickname, u.about_me, upi.path, ups.private_account from users u 
	JOIN user_profile_images upi ON upi.user_id = u.id 
	JOIN user_privacy_settings ups ON ups.id = u.id 
	where u.id  = $1;`
	err := m.DB.QueryRow(query, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.Nickname, &user.AboutMe, &user.Avatar, &user.Private)

	user.Avatar = config.AVATAR_PATH_URL + user.Avatar

	return user, err
}

// GetUserFollowers gets all user followers
func (m *sqliteDBRepo) GetUserFollowers(id int) ([]models.Follow, error) {
	var followers []models.Follow

	query := `SELECT u.id, u.first_name, u.last_name, (SELECT COUNT(*) FROM followers WHERE user_id = f.follower_id), upi.path FROM followers f 
	JOIN users u ON u.id = f.follower_id
	JOIN user_profile_images upi ON upi.user_id = f.follower_id
	WHERE f.user_id = $1;`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return followers, err
	}

	for rows.Next() {
		var follower models.Follow

		if rows.Scan(&follower.Id, &follower.FirstName, &follower.LastName, &follower.Followers, &follower.Avatar) != nil {
			return followers, err
		}

		follower.Type = "follower"
		follower.Avatar = config.AVATAR_PATH_URL + follower.Avatar

		followers = append(followers, follower)
	}

	return followers, nil
}

// GetUserFollows gets all user follows
func (m *sqliteDBRepo) GetUserFollows(id int) ([]models.Follow, error) {
	var follows []models.Follow

	query := `SELECT u.id, u.first_name, u.last_name, (SELECT COUNT(*) FROM followers WHERE user_id = f.follower_id), upi.path FROM followers f 
	JOIN users u ON u.id = f.user_id
	JOIN user_profile_images upi ON upi.user_id = f.user_id
	WHERE f.follower_id = $1;`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return follows, err
	}

	for rows.Next() {
		var follow models.Follow

		if rows.Scan(&follow.Id, &follow.FirstName, &follow.LastName, &follow.Followers, &follow.Avatar) != nil {
			return follows, err
		}

		follow.Type = "following"
		follow.Avatar = config.AVATAR_PATH_URL + follow.Avatar

		follows = append(follows, follow)
	}

	return follows, nil
}

// GetUserAvatar gets user's avatar path
func (m *sqliteDBRepo) GetUserAvatar(id int) (string, error) {
	var result string

	query := `select upi.path from user_profile_images upi where upi.user_id = $1`
	err := m.DB.QueryRow(query, id).Scan(&result)

	return result, err
}

// CheckEmailExistence checks if email is already taken
func (m *sqliteDBRepo) CheckEmailExistence(email string) (int, error) {
	var id int

	query := `select id from users where email = $1`
	err := m.DB.QueryRow(query, email).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return id, nil
}

// CheckSessionExistence checks if session token exists in the database
func (m *sqliteDBRepo) CheckSessionExistence(token string) (int, error) {
	var id int

	query := `select user_id from user_sessions where session = $1`
	err := m.DB.QueryRow(query, token).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return id, nil
}

// RemoveEventParticipatingRequest removes a event request of certain user
func (m *sqliteDBRepo) RemoveEventParticipatingRequest(eid, uid int) error {
	query := `DELETE FROM group_event_request WHERE user_id = $1 AND event_id = $2`
	_, err := m.DB.Exec(query, uid, eid)

	return err
}

// RemoveFollowRequest removes a follow request of certain user
func (m *sqliteDBRepo) RemoveFollowRequest(sourceId, destId int) error {
	query := `DELETE FROM follow_requests WHERE follow_from = $1 AND follow_to = $2`
	_, err := m.DB.Exec(query, sourceId, destId)

	return err
}

// UpdateSessionToken updates token to a new one for user
func (m *sqliteDBRepo) UpdateSessionToken(token string, id int) error {
	query := `update user_sessions set session = $1 where user_id = $2`
	_, err := m.DB.Exec(query, token, id)

	return err
}

// UpdateUserAvatar updates user's avatar
func (m *sqliteDBRepo) UpdateUserAvatar(id int, path string) error {
	query := `update user_profile_images set path = $1 where user_id = $2`
	_, err := m.DB.Exec(query, path, id)

	return err
}

// UpdateUserProfile updates user profile data
func (m *sqliteDBRepo) UpdateUserProfile(id int, aboutMe, nickname string) error {
	query := `update users set about_me = $1, nickname = $2 where id = $3`
	_, err := m.DB.Exec(query, aboutMe, nickname, id)

	return err
}

// UpdateUserPrivacy updates user privacy data
func (m *sqliteDBRepo) UpdateUserPrivacy(id int, private bool) error {
	query := `update user_privacy_settings set private_account = $1 where user_id = $2`
	_, err := m.DB.Exec(query, private, id)

	return err
}

// GetUserHash gets user's password hash for further compare
func (m *sqliteDBRepo) GetUserHash(id int) (string, error) {
	var hash string

	query := `select password from users where id = $1`
	err := m.DB.QueryRow(query, &id).Scan(&hash)
	if err != nil {
		return "", err
	}

	return hash, nil
}
