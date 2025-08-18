-- +goose Up
CREATE TABLE feeds(
  id UUID PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  url TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL,
  user_id UUID REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE feeds;