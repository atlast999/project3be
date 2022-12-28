-- name: GetMyList :many
SELECT web_apps.* FROM web_apps, my_lists
WHERE my_lists.user_id = $1
AND my_lists.app_id = web_apps.id
OFFSET $2
LIMIT $3;

-- name: AddMyList :exec
INSERT INTO my_lists (user_id, app_id)
VALUES ($1, $2);
