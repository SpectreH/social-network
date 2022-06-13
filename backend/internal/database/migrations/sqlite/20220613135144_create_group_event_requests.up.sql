CREATE TABLE group_event_request (
	id INTEGER PRIMARY KEY,
  request_status_id INTEGER NOT NULL,
  group_id INTEGER NOT NULL,
  event_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  requested_at DATE NOT NULL,
  FOREIGN KEY(request_status_id) REFERENCES follow_request_statuses(id),
  FOREIGN KEY(group_id) REFERENCES groups(id),
  FOREIGN KEY(event_id) REFERENCES group_events(id),
  FOREIGN KEY(user_id) REFERENCES users(id)
);