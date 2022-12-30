package test

import (
	"testing"

	db "github.com/atlast999/project3be/db/gen"
	"github.com/atlast999/project3be/helper"
	"github.com/atlast999/project3be/repository"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) db.User {
	var username = helper.RandString(6)
	var password = helper.RandString(6)
	user, err := repository.CreateUser(db.CreateUserParams{
		Username: username,
		Password: password,
	}, txInstance)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, username, user.Username)
	return user
}

func TestCreateUser(t *testing.T) {
	user := createRandomUser(t)
	require.NotNil(t, user.ID)
}

func TestGetUser(t *testing.T) {
	username := helper.RandString(6)
	password := helper.RandString(8)
	hashed, err := helper.GeneratePassword(password)
	require.NoError(t, err)
	user, err := repository.CreateUser(db.CreateUserParams{
		Username: username,
		Password: hashed,
	}, txInstance)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	foundUser, err := repository.GetUser(user.Username, txInstance)
	require.NoError(t, err)
	require.NotEmpty(t, foundUser)
	err = helper.CheckPassword(password, foundUser.Password)
	require.NoError(t, err)
	require.Equal(t, user.ID, foundUser.ID)
}

func createRandomWebAppParams() db.CreateWebAppParams {
	return db.CreateWebAppParams{
		Name:  helper.RandString(6),
		Url:   helper.RandString(6),
		Image: helper.RandString(6),
	}
}

func TestCreateWebApp(t *testing.T) {
	user := createRandomUser(t)
	params := createRandomWebAppParams()
	webApp, err := repository.CreateWebApp(int(user.ID), params, txInstance)
	require.NoError(t, err)
	require.NotEmpty(t, webApp)
	require.Equal(t, params.Name, webApp.Name)
	require.Equal(t, webApp.CollectionID, 0)
}

func TestGetMyList(t *testing.T) {
	user := createRandomUser(t)
	size := 5
	for i := 0; i < size; i++ {
		params := createRandomWebAppParams()
		repository.CreateWebApp(int(user.ID), params, txInstance)
	}
	webApps, err := repository.GetMyList(db.GetMyListParams{
		UserID: user.ID,
		Offset: 0,
		Limit:  5,
	}, txInstance)
	require.NoError(t, err)
	require.Equal(t, len(webApps), size)
}

func TestShareMyList(t *testing.T) {
	user := createRandomUser(t)
	size := 5
	for i := 0; i < size; i++ {
		params := createRandomWebAppParams()
		repository.CreateWebApp(int(user.ID), params, txInstance)
	}
	var collectionName = helper.RandString(6)
	collection, err := repository.ShareMyList(int(user.ID), collectionName, dbInstance)
	require.NoError(t, err)
	require.NotEmpty(t, collection)
	require.Equal(t, collectionName, collection.Name)
	require.Equal(t, user.ID, collection.OwnerID)
}

func TestGetCollections(t *testing.T) {
	user := createRandomUser(t)
	params := createRandomWebAppParams()
	repository.CreateWebApp(int(user.ID), params, txInstance)
	var collectionName = helper.RandString(6)
	repository.ShareMyList(int(user.ID), collectionName, dbInstance)
	collections, err := repository.GetCollections(dbQueries)
	require.NoError(t, err)
	require.NotEmpty(t, collections)
}

func TestGetCollectionDetail(t *testing.T) {
	user := createRandomUser(t)
	size := 5
	for i := 0; i < size; i++ {
		params := createRandomWebAppParams()
		repository.CreateWebApp(int(user.ID), params, txInstance)
	}
	var collectionName = helper.RandString(6)
	collection, _ := repository.ShareMyList(int(user.ID), collectionName, dbInstance)
	webApps, err := repository.GetCollectionDetail(int(collection.ID), dbQueries)
	require.NoError(t, err)
	require.Equal(t, len(webApps), size)
}

func TestTakeCollection(t *testing.T) {
	sharingUser := createRandomUser(t)
	size := 5
	for i := 0; i < size; i++ {
		params := createRandomWebAppParams()
		repository.CreateWebApp(int(sharingUser.ID), params, txInstance)
	}
	var collectionName = helper.RandString(6)
	collection, _ := repository.ShareMyList(int(sharingUser.ID), collectionName, dbInstance)
	sharedApps, _ := repository.GetMyList(db.GetMyListParams{
		UserID: sharingUser.ID,
		Offset: 0,
		Limit:  5,
	}, txInstance)
	takingUser := createRandomUser(t)
	err := repository.TakeCollection(int(takingUser.ID), int(collection.ID), dbInstance)
	takenApps, _ := repository.GetMyList(db.GetMyListParams{
		UserID: takingUser.ID,
		Offset: 0,
		Limit:  5,
	}, txInstance)
	require.NoError(t, err)
	require.Equal(t, len(takenApps), len(sharedApps))
}
