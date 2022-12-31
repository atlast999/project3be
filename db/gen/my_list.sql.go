// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: my_list.sql

package db

import (
	"context"
)

const addMyList = `-- name: AddMyList :exec
INSERT INTO my_lists (user_id, app_id)
VALUES ($1, $2)
`

type AddMyListParams struct {
	UserID int32 `json:"user_id"`
	AppID  int32 `json:"app_id"`
}

func (q *Queries) AddMyList(ctx context.Context, arg AddMyListParams) error {
	_, err := q.db.ExecContext(ctx, addMyList, arg.UserID, arg.AppID)
	return err
}

const getMyList = `-- name: GetMyList :many
SELECT web_apps.id, web_apps.name, web_apps.url, web_apps.image, web_apps.collection_id FROM web_apps, my_lists
WHERE my_lists.user_id = $1
AND my_lists.app_id = web_apps.id
OFFSET $2
LIMIT $3
`

type GetMyListParams struct {
	UserID int32 `json:"user_id"`
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) GetMyList(ctx context.Context, arg GetMyListParams) ([]WebApp, error) {
	rows, err := q.db.QueryContext(ctx, getMyList, arg.UserID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []WebApp{}
	for rows.Next() {
		var i WebApp
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.Image,
			&i.CollectionID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeMyList = `-- name: RemoveMyList :exec
WITH deleted_app_ids AS (
    DELETE FROM my_lists
    WHERE user_id = $1
    RETURNING app_id
)
DELETE FROM web_apps
WHERE id IN (
    SELECT app_id FROM deleted_app_ids
)
`

func (q *Queries) RemoveMyList(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, removeMyList, userID)
	return err
}
