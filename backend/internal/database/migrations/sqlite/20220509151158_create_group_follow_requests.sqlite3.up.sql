CREATE TABLE group_follow_requests (
	id INTEGER PRIMARY KEY,
  request_status_id INTEGER NOT NULL,
  group_id INTEGER NOT NULL,
  creator_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  invite BOOLEAN NOT NULL,
  requested_at DATE NOT NULL,
  FOREIGN KEY(request_status_id) REFERENCES follow_request_statuses(id),
  FOREIGN KEY(group_id) REFERENCES groups(id),
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(creator_id) REFERENCES users(id)
);