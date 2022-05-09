CREATE TABLE post_images (
	id INTEGER PRIMARY KEY,
  post_id INTEGER NOT NULL,
  path TEXT NULL DEFAULT "-", 
  FOREIGN KEY(post_id) REFERENCES posts(id)
);