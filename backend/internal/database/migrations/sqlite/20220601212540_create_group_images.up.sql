CREATE TABLE group_images (
	id INTEGER PRIMARY KEY,
  group_id INTEGER NOT NULL,
  path TEXT NULL DEFAULT "-", 
  FOREIGN KEY(group_id) REFERENCES "groups"(id)
);