-- name: GetMyList :many
SELECT web_apps.* FROM web_apps, my_lists
WHERE my_lists.user_id = $1
AND my_lists.app_id = web_apps.id
OFFSET $2
LIMIT $3;

-- name: AddMyList :exec
INSERT INTO my_lists (user_id, app_id)
VALUES ($1, $2);

-- name: RemoveMyList :exec
WITH deleted_app_ids AS (
    DELETE FROM my_lists
    WHERE user_id = $1
    RETURNING app_id
)
DELETE FROM web_apps
WHERE id IN (
    SELECT app_id FROM deleted_app_ids
);