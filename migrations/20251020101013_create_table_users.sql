-- +goose Up
CREATE TABLE users (
    id SERIAL primary key,
    username VARCHAR(50) UNIQUE,
    email VARCHAR(100) UNIQUE,
    password varchar(50)
);


-- +goose Down
DROP TABLE users;
