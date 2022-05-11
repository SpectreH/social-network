CREATE TABLE user_sessions (
	id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  session TEXT NOT NULL DEFAULT "-",
  FOREIGN KEY(user_id) REFERENCES user(id)
);