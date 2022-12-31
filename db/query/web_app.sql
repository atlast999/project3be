-- name: CreateWebApp :one
INSERT INTO web_apps (name, url, image)
VALUES ($1, $2, $3)
RETURNING *;

-- name: AddToCollection :exec
INSERT INTO web_apps (name, url, image, collection_id)
SELECT name, url, image, @collection_id::int FROM web_apps
WHERE id IN (SELECT app_id FROM my_lists WHERE user_id = $1);

-- name: GetByCollection :many
SELECT * FROM web_apps
WHERE collection_id = @collection_id::int
OFFSET $1
LIMIT $2;

-- name: TakeCollection :exec
WITH new_app_ids AS (
    INSERT INTO web_apps (name, url, image)
    SELECT name, url, image FROM web_apps
    WHERE collection_id = @collection_id::int
    RETURNING id
)
INSERT INTO my_lists (user_id, app_id)
SELECT @user_id::int, id FROM new_app_ids;
