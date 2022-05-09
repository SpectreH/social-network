CREATE TABLE group_events (
	id INTEGER PRIMARY KEY,
  group_id INTEGER NOT NULL,
  title TEXT NOT NULL, 
  description BLOB NOT NULL, 
  date DATE NOT NULL,
  created_at DATE NOT NULL, 
  FOREIGN KEY(group_id) REFERENCES groups(id)
);