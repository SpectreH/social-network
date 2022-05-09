CREATE TABLE user_profile_images (
	id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  path TEXT NOT NULL DEFAULT "./img/default_picture.png", 
  FOREIGN KEY(user_id) REFERENCES users(id)
);