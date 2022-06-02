CREATE TABLE posts (
	id INTEGER PRIMARY KEY,
  group_id INTEGER DEFAULT 0,
  user_id INTEGER NOT NULL,
  share_id INTEGER NOT NULL,
	content BLOB NOT NULL,
  created_at DATE NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(share_id) REFERENCES share_types(id)
);