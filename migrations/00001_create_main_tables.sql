-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id uuid primary key default gen_random_uuid(),
    first_name VARCHAR(255),
    username VARCHAR(255)
);

CREATE TABLE posts (
    id uuid primary key default gen_random_uuid(),
    user_id uuid REFERENCES users(id),
    title VARCHAR(255),
    body TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE posts;
-- +goose StatementEnd
