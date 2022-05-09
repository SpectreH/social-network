CREATE TABLE follow_requests (
	id INTEGER PRIMARY KEY,
  request_status_id INTEGER NOT NULL,
  follow_from INTEGER NOT NULL,
  follow_to INTEGER NOT NULL,
  requested_at DATE NOT NULL,
  FOREIGN KEY(follow_from) REFERENCES users(id),
  FOREIGN KEY(follow_to) REFERENCES users(id),
  FOREIGN KEY(request_status_id) REFERENCES follow_request_statuses(id)
);