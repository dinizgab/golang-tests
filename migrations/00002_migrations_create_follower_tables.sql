-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS followers (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid references users(id),
  follower_id uuid references users(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS followers;
-- +goose StatementEnd
