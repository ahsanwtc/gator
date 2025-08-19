-- +goose Up
CREATE TABLE feed_follows (
  id UUID PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  user_id UUID REFERENCES users (id) ON DELETE CASCADE NOT NULL,
  feed_id UUID REFERENCES feeds (id) ON DELETE CASCADE NOT NULL,
  UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;