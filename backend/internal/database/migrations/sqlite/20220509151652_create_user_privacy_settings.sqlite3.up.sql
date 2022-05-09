CREATE TABLE user_privacy_settings (
	id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  private_account BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY(user_id) REFERENCES users(id)
);