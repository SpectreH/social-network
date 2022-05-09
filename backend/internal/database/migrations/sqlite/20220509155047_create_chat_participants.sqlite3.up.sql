CREATE TABLE chat_participants (
	id INTEGER PRIMARY KEY,
  chat_id INTEGER NOT NULL,
  participant_id INTEGER NOT NULL,L,
  FOREIGN KEY(chat_id) REFERENCES chats(id),
  FOREIGN KEY(participant_id) REFERENCES users(id)
);