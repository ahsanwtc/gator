-- +goose Up
CREATE TABLE posts(
  id UUID PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  title VARCHAR(255) NOT NULL,
  url TEXT NOT NULL,
  description TEXT NOT NULL,
  published_at TIMESTAMPTZ NULL,
  feed_id UUID REFERENCES feeds (id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE posts;