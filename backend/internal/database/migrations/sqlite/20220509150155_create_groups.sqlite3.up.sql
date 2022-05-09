CREATE TABLE "groups" (
	id INTEGER PRIMARY KEY,
  chat_id INTEGER NOT NULL,
  creator_id INTEGER NOT NULL,
  title TEXT NOT NULL, 
  description BLOB NOT NULL, 
  private BOOLEAN NOT NULL,
  created_at DATE NOT NULL, 
  FOREIGN KEY(creator_id) REFERENCES users(id),
  FOREIGN KEY(chat_id) REFERENCES chats(id)
);