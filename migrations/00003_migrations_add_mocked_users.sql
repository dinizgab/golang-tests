-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO users (id, first_name, username) VALUES ('30d4971f-0677-4b99-a03a-cd8f17c8d893', 'John', 'john_doe');
INSERT INTO users (id, first_name, username) VALUES ('f15c6c35-76b8-4d93-8971-727d99df6b4d', 'Jane', 'jane_doe');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users
WHERE id = '30d4971f-0677-4b99-a03a-cd8f17c8d893'
    OR id = 'f15c6c35-76b8-4d93-8971-727d99df6b4d';
-- +goose StatementEnd
