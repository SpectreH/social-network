CREATE TABLE group_event_participants (
	id INTEGER PRIMARY KEY,
  participant_id INTEGER NOT NULL,
  event_id INTEGER NOT NULL,
  will_attend BOOLEAN NOT NULL,
  FOREIGN KEY(participant_id) REFERENCES group_membership(user_id),
  FOREIGN KEY(event_id) REFERENCES group_events(id)
);