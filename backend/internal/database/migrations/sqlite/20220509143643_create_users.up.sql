CREATE TABLE users (
	id INTEGER PRIMARY KEY,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	birth_date TEXT NOT NULL,
	nickname TEXT NOT NULL DEFAULT "-",
	about_me TEXT NOT NULL DEFAULT "-",
  password TEXT NOT NULL
);