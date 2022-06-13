CREATE TABLE group_event_images (
	id INTEGER PRIMARY KEY,
  event_id INTEGER NOT NULL,
  path TEXT NULL DEFAULT "-", 
  FOREIGN KEY(event_id) REFERENCES group_events(id)
);