CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(40) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    display_name VARCHAR(40) NOT NULL DEFAULT '',
    bio VARCHAR(2000) NOT NULL DEFAULT '',
    color VARCHAR(7)
);

CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    name VARCHAR(50) NOT NULL,
    url VARCHAR(100) NOT NULL,
    color VARCHAR(7)
);
