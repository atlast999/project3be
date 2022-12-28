package repository

import (
	"context"

	db "github.com/atlast999/project3be/db/gen"
)

func CreateUser(username, password string, queries *db.Queries) (db.User, error) {
	return queries.CreateUser(context.Background(), db.CreateUserParams{
		Username: username,
		Password: password,
	})
}

func GetUser(username, password string, queries *db.Queries) (db.User, error) {
	return queries.GetUser(context.Background(), db.GetUserParams{
		Username: username,
		Password: password,
	})
}

type WebAppParams struct {
	UserID int
	Name string
	Url string
	Image string
}

func CreateWebApp(param WebAppParams, queries *db.Queries) (db.WebApp, error) {
	webApp, err := queries.CreateWebApp(context.Background(), db.CreateWebAppParams{
		Name: param.Name,
		Url: param.Url,
		Image: param.Image,
	})
	if err != nil {
		return webApp, err
	}
	err = queries.AddMyList(context.Background(), db.AddMyListParams{
		UserID: int32(param.UserID),
		AppID: webApp.ID,
	})
	return webApp, err
}

func GetMyList(userID int, queries *db.Queries) ([]db.WebApp, error) {
	return queries.GetMyList(context.Background(), db.GetMyListParams{
		UserID: int32(userID),
		Offset: 0,
		Limit: 10,
	})
}

func ShareMyList(userID int, collectionName string, queries *db.Queries) (db.Collection, error) {
	collection, err := queries.CreateCollection(context.Background(), db.CreateCollectionParams{
		Name: collectionName,
		OwnerID: int32(userID),
	})
	if err != nil {
		return collection, err
	}
	err = queries.AddToCollection(context.Background(), db.AddToCollectionParams{
		UserID: int32(userID),
		CollectionID: collection.ID,
	})
	return collection, err
}

