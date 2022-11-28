CREATE TABLE users (
	user_id serial PRIMARY KEY,
	email TEXT UNIQUE,
	password TEXT NOT NULL
);

CREATE TABLE links (
	link_id serial PRIMARY KEY,
	target_url TEXT  NOT NULL,
	shortened_url TEXT UNIQUE,
	created_on TIMESTAMP NOT NULL,
	clicks INT DEFAULT 0,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users (user_id)  ON DELETE CASCADE
);
