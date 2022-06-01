package sqlite

import (
	"database/sql"
	"fmt"
	"social-network/internal/config"
	"social-network/internal/models"
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

// InsertUserFollowRequest inserts uses's follow request with status pending
func (m *sqliteDBRepo) InsertUserFollowRequest(srcId, targetId int) error {
	query := `insert into follow_requests (request_status_id, follow_from, follow_to, requested_at) values ($1, $2, $3, $4);`
	_, err := m.DB.Exec(query, 1, srcId, targetId, time.Now())

	return err
}

// InsertPost inserts new post
func (m *sqliteDBRepo) InsertPost(post models.Post) (int, error) {
	var id int

	fmt.Println(post.Content, post.CreatedAt)

	query := `insert into posts (user_id, share_id, content, created_at) values ($1, $2, $3, $4) returning id;`
	err := m.DB.QueryRow(query, post.AuthId, post.ShareId, post.Content, post.CreatedAt).Scan(&id)

	return id, err
}

// InsertPostPicture inserts post picture
func (m *sqliteDBRepo) InsertPostPicture(id int, path string) error {
	query := `insert into post_images (post_id, path) values ($1, $2);`
	_, err := m.DB.Exec(query, id, path)

	return err
}

// CheckFollowRequest checks if follow request already exists
func (m *sqliteDBRepo) CheckFollowRequest(srcId, targetId int) (int, error) {
	var res int
	query := `select COUNT(*) from follow_requests WHERE follow_from = $1 AND follow_to= $2;`
	err := m.DB.QueryRow(query, srcId, targetId).Scan(&res)

	return res, err
}

// CheckProfileIsPivate chekcs if user has private account
func (m *sqliteDBRepo) CheckProfileIsPivate(id int) (bool, error) {
	var res bool

	query := `select ups.private_account from user_privacy_settings ups where ups.user_id = $1;`
	err := m.DB.QueryRow(query, id).Scan(&res)

	return res, err
}

// CheckAlreadyFollowed chekcs if user already followed certain user
func (m *sqliteDBRepo) CheckAlreadyFollowed(srcId, targetId int) (int, error) {
	var res int
	query := `select COUNT(*) from followers WHERE follower_id = $1 AND user_id = $2;`
	err := m.DB.QueryRow(query, srcId, targetId).Scan(&res)

	return res, err
}

// FollowUser makes a record with follow
func (m *sqliteDBRepo) FollowUser(srcId, targetId int) error {
	query := `insert into followers (user_id, follower_id, followed_at) values ($1, $2, $3);`
	_, err := m.DB.Exec(query, targetId, srcId, time.Now())

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

	query := `SELECT p.id, p.user_id, u.first_name, u.last_name, ufi.path, p.share_id, p.content, pi.path, p.created_at FROM posts p
	JOIn post_images pi ON pi.post_id = p.id
	JOIn users u ON u.id = p.user_id
	JOIn user_profile_images ufi ON ufi.user_id = p.user_id
	WHEre p.id = $1;`
	err := m.DB.QueryRow(query, id).Scan(&res.Id, &res.AuthId, &res.FirstName, &res.LastName, &res.Avatar, &res.ShareId, &res.Content, &res.Picture, &res.CreatedAt)

	res.Picture = config.AVATAR_PATH_URL + res.Picture
	res.Avatar = config.AVATAR_PATH_URL + res.Avatar

	return res, err
}

// GetPostComments gets all post comments
func (m *sqliteDBRepo) GetPostComments(id int) ([]models.Comment, error) {
	var comments []models.Comment

	query := `SELECT c.id, c.user_id, u.first_name, u.last_name, ufi.path, c.content, c.created_at, ci.path FROM comments c
	JOIN comment_images ci ON ci.comment_id= c.id
	JOIN users u ON u.id = c.user_id  
	JOIn user_profile_images ufi ON ufi.user_id = c.user_id
	WHEre c.post_id = $1;`

	rows, err := m.DB.Query(query, id)
	if err != nil && err != sql.ErrNoRows {
		return comments, err
	}

	for rows.Next() {
		var comment models.Comment

		if rows.Scan(&comment.Id, &comment.AuthId, &comment.FirstName, &comment.LastName, &comment.Avatar, &comment.Content, &comment.CreatedAt, &comment.Picture) != nil {
			return comments, err
		}

		comment.Picture = config.AVATAR_PATH_URL + comment.Picture
		comment.Avatar = config.AVATAR_PATH_URL + comment.Avatar

		comments = append(comments, comment)
	}

	return comments, nil
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
