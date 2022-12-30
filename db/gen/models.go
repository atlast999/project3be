// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import ()

type Collection struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	OwnerID int32  `json:"owner_id"`
}

type MyList struct {
	UserID int32 `json:"user_id"`
	AppID  int32 `json:"app_id"`
}

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type WebApp struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Url          string `json:"url"`
	Image        string `json:"image"`
	CollectionID int32  `json:"collection_id"`
}
