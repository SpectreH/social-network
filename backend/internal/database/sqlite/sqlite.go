package sqlite

import (
	"database/sql"
	"social-network/internal/models"
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
	query := `update user_sessions set sessio = $1 where user_id = $2`
	_, err := m.DB.Exec(query, token, id)

	return err
}
