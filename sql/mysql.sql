CREATE DATABASE IF NOT EXISTS shorgot;

USE shorgot;

CREATE TABLE IF NOT EXISTS users (
	user_id INT auto_increment PRIMARY KEY,
	email VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(150) NOT NULL
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS links (
	link_id INT auto_increment PRIMARY KEY,
	target_url VARCHAR ( 255 )  NOT NULL,
	shortened_url VARCHAR ( 150 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
	clicks INT default 0,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(user_id)
) ENGINE=INNODB;


