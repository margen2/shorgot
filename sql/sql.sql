CREATE TABLE users (
	user_id serial PRIMARY KEY,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 150 ) NOT NULL
);

CREATE TABLE links (
	link_id serial PRIMARY KEY,
	target_url VARCHAR ( 255 )  NOT NULL,
	shortened_url VARCHAR ( 150 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id)
      REFERENCES users (user_id),
	clicks INT NOT NULL
);
