CREATE TABLE chat_messages (
	id INTEGER PRIMARY KEY,
  chat_id INTEGER NOT NULL,
  author_id INTEGER NOT NULL,
  content BLOB NOT NULL,
  created_at DATE NOT NULL,
  FOREIGN KEY(chat_id) REFERENCES chats(id),
  FOREIGN KEY(author_id) REFERENCES users(id)
);