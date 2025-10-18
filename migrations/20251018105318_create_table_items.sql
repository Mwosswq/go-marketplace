-- +goose Up
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50),
    description VARCHAR(2000),
    created_at TIMESTAMP DEFAULT NOW(),
    price FLOAT
);

-- +goose Down
DROP TABLE items;

