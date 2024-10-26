CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(40) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    display_name VARCHAR(40),
    bio TEXT
);
