CREATE TABLE comment_images (
	id INTEGER PRIMARY KEY,
  comment_id INTEGER NOT NULL,
  path TEXT NULL DEFAULT "-", 
  FOREIGN KEY(comment_id) REFERENCES comments(id)
);