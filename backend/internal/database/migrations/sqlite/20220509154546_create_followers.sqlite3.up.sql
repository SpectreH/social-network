CREATE TABLE followers (
	id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  follower_id INTEGER NOT NULL,
  followed_at DATE NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(follower_id) REFERENCES users(id)
);