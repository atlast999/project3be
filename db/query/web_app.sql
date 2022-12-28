-- name: CreateWebApp :one
INSERT INTO web_apps (name, url, image)
VALUES ($1, $2, $3)
RETURNING *;

-- name: AddToCollection :exec
UPDATE web_apps
SET collection_id = @collection_id::int
WHERE id IN (SELECT app_id FROM my_lists WHERE user_id = $1);

-- name: GetByCollection :many
SELECT * FROM web_apps
WHERE collection_id = $1
OFFSET $2
LIMIT $3;