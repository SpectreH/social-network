package sqlite

import (
	"database/sql"
	"fmt"
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

	return profile, err
}

// InsertUserFollowRequest inserts uses's follow request with status pending
func (m *sqliteDBRepo) InsertUserFollowRequest(srcId, targetId int) error {
	query := `insert into follow_requests (request_status_id, follow_from, follow_to, requested_at) values ($1, $2, $3, $4);`
	_, err := m.DB.Exec(query, 1, srcId, targetId, time.Now())

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

// GetUserData gets information about user
func (m *sqliteDBRepo) GetUserData(id int) (models.User, error) {
	var user models.User

	query := `select u.id, u.first_name, u.last_name, u.email, u.birth_date, u.nickname, u.about_me, upi.path, ups.private_account from users u 
	JOIN user_profile_images upi ON upi.user_id = u.id 
	JOIN user_privacy_settings ups ON ups.id = u.id 
	where u.id = $1;`
	err := m.DB.QueryRow(query, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.Nickname, &user.AboutMe, &user.Avatar, &user.Private)

	return user, err
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
