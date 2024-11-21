CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(40) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    display_name VARCHAR(40) NOT NULL DEFAULT '',
    bio VARCHAR(2000) NOT NULL DEFAULT ''
);

CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    name varchar(50) NOT NULL,
    url varchar(100) NOT NULL
);
