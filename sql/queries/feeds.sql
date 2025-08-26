-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET created_at = now(), last_fetched_at = now()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * from feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;