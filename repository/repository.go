package repository

import (
	"context"
	"database/sql"

	db "github.com/atlast999/project3be/db/gen"
	"github.com/atlast999/project3be/db/transaction"
)

func CreateUser(params db.CreateUserParams, txInstance *transaction.TxInstance) (db.User, error) {
	return txInstance.Queries.CreateUser(context.Background(), params)
}

func GetUser(params db.GetUserParams, txInstance *transaction.TxInstance) (db.User, error) {
	return txInstance.Queries.GetUser(context.Background(), params)
}

func CreateWebApp(userId int, params db.CreateWebAppParams, txInstance *transaction.TxInstance) (db.WebApp, error) {
	var webApp db.WebApp
	txErr := txInstance.ExecTransaction(context.Background(), func(queries *db.Queries) error {
		var err error
		webApp, err = queries.CreateWebApp(context.Background(), params)
		if err != nil {
			return err
		}
		err = queries.AddMyList(context.Background(), db.AddMyListParams{
			UserID: int32(userId),
			AppID:  webApp.ID,
		})
		return err
	})
	return webApp, txErr
}

func GetMyList(userID int, queries *db.Queries) ([]db.WebApp, error) {
	return queries.GetMyList(context.Background(), db.GetMyListParams{
		UserID: int32(userID),
		Offset: 0,
		Limit:  10,
	})
}

func ShareMyList(userID int, collectionName string, dbInstance *sql.DB) (db.Collection, error) {
	txInstance := transaction.NewTxInstance(dbInstance)
	var collection db.Collection
	txErr := txInstance.ExecTransaction(context.Background(), func(queries *db.Queries) error {
		var err error
		collection, err = queries.CreateCollection(context.Background(), db.CreateCollectionParams{
			Name:    collectionName,
			OwnerID: int32(userID),
		})
		if err != nil {
			return err
		}
		err = queries.AddToCollection(context.Background(), db.AddToCollectionParams{
			UserID:       int32(userID),
			CollectionID: collection.ID,
		})
		return err
	})
	return collection, txErr
}

func GetCollections(queries *db.Queries) ([]db.Collection, error) {
	return queries.GetCollections(context.Background(), db.GetCollectionsParams{
		Offset: 0,
		Limit:  10,
	})
}

func GetCollectionDetail(collectionId int, queries *db.Queries) ([]db.WebApp, error) {
	return queries.GetByCollection(context.Background(), db.GetByCollectionParams{
		CollectionID: int32(collectionId),
		Offset:       0,
		Limit:        10,
	})
}

func TakeCollection(userId, collectionId int, dbInstance *sql.DB) error {
	txInstance := transaction.NewTxInstance(dbInstance)
	return txInstance.ExecTransaction(context.Background(), func(queries *db.Queries) error {
		err := queries.RemoveMyList(context.Background(), int32(userId))
		if err != nil {
			return err
		}
		return queries.TakeCollection(context.Background(), db.TakeCollectionParams{
			UserID:       int32(userId),
			CollectionID: int32(collectionId),
		})
	})
}
