-- name: CreateCollection :one
INSERT INTO collections (name, owner_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetCollections :many
SELECT * FROM collections
OFFSET $1
LIMIT $2;
