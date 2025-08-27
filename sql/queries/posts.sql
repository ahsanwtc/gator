-- name: CreatePost :one
INSERT INTO posts (id, title, url, description, feed_id, published_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT p.id, p.title, p.url, p.description, p.updated_at, f.id AS feed_id, f.name AS feed_name
FROM feed_follows AS ff
INNER JOIN feeds AS f ON f.id = ff.feed_id
INNER JOIN posts AS p ON p.feed_id = f.id
WHERE ff.user_id = $1
ORDER BY COALESCE(p.published_at, p.updated_at) DESC
LIMIT $2;