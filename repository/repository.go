package repository

import (
	"context"

	db "github.com/atlast999/project3be/db/gen"
	"github.com/atlast999/project3be/db/transaction"
)

func CreateUser(params db.CreateUserParams, txInstance *transaction.TxInstance) (db.User, error) {
	return txInstance.Queries.CreateUser(context.Background(), params)
}

func GetUser(username string, txInstance *transaction.TxInstance) (db.User, error) {
	return txInstance.Queries.GetUser(context.Background(), username)
}

func CreateWebApp(userId int32, params db.CreateWebAppParams, txInstance *transaction.TxInstance) (db.WebApp, error) {
	var webApp db.WebApp
	txErr := txInstance.ExecTransaction(context.Background(), func(queries *db.Queries) error {
		var err error
		webApp, err = queries.CreateWebApp(context.Background(), params)
		if err != nil {
			return err
		}
		err = queries.AddMyList(context.Background(), db.AddMyListParams{
			UserID: userId,
			AppID:  webApp.ID,
		})
		return err
	})
	return webApp, txErr
}

func GetMyList(params db.GetMyListParams, txInstance *transaction.TxInstance) ([]db.WebApp, error) {
	return txInstance.Queries.GetMyList(context.Background(), params)
}

func ShareMyList(userID int32, collectionName string, txInstance *transaction.TxInstance) (db.Collection, error) {
	var collection db.Collection
	txErr := txInstance.ExecTransaction(context.Background(), func(queries *db.Queries) error {
		var err error
		collection, err = queries.CreateCollection(context.Background(), db.CreateCollectionParams{
			Name:    collectionName,
			OwnerID: userID,
		})
		if err != nil {
			return err
		}
		err = queries.AddToCollection(context.Background(), db.AddToCollectionParams{
			UserID:       userID,
			CollectionID: collection.ID,
		})
		return err
	})
	return collection, txErr
}

func GetCollections(params db.GetCollectionsParams, txInstance *transaction.TxInstance) ([]db.Collection, error) {
	return txInstance.Queries.GetCollections(context.Background(), params)
}

func GetCollectionDetail(params db.GetByCollectionParams, txInstance *transaction.TxInstance) ([]db.WebApp, error) {
	return txInstance.Queries.GetByCollection(context.Background(), params)
}

func TakeCollection(userId, collectionId int32, txInstance *transaction.TxInstance) error {
	return txInstance.ExecTransaction(context.Background(), func(queries *db.Queries) error {
		err := queries.RemoveMyList(context.Background(), userId)
		if err != nil {
			return err
		}
		return queries.TakeCollection(context.Background(), db.TakeCollectionParams{
			UserID:       userId,
			CollectionID: collectionId,
		})
	})
}
