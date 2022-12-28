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
	user, err := repository.CreateUser(username, password, dbQueries)
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
	user := createRandomUser(t)
	foundUser, err := repository.GetUser(user.Username, user.Password, dbQueries)
	require.NoError(t, err)
	require.NotEmpty(t, foundUser)
	require.Equal(t, user.ID, foundUser.ID)
}

func createRandomWebAppParams(userID int) repository.WebAppParams {
	return repository.WebAppParams{
		UserID: userID,
		Name:   helper.RandString(6),
		Url:    helper.RandString(6),
		Image:  helper.RandString(6),
	}
}

func TestCreateWebApp(t *testing.T) {
	user := createRandomUser(t)
	params := createRandomWebAppParams(int(user.ID))
	webApp, err := repository.CreateWebApp(
		params,
		dbQueries,
	)
	require.NoError(t, err)
	require.NotEmpty(t, webApp)
	require.Equal(t, params.Name, webApp.Name)
	require.False(t, webApp.CollectionID.Valid)
}

func TestGetMyList(t *testing.T) {
	user := createRandomUser(t)
	size := 5
	for i := 0; i < size; i++ {
		params := createRandomWebAppParams(int(user.ID))
		repository.CreateWebApp(params, dbQueries)
	}
	webApps, err := repository.GetMyList(int(user.ID), dbQueries)
	require.NoError(t, err)
	require.Equal(t, len(webApps), size)
}

func TestShareMyList(t *testing.T) {
	user := createRandomUser(t)
	size := 5
	for i := 0; i < size; i++ {
		params := createRandomWebAppParams(int(user.ID))
		repository.CreateWebApp(params, dbQueries)
	}
	var collectionName = helper.RandString(6)
	collection, err := repository.ShareMyList(int(user.ID), collectionName, dbQueries)
	require.NoError(t, err)
	require.NotEmpty(t, collection)
	require.Equal(t, collectionName, collection.Name)
	require.Equal(t, user.ID, collection.OwnerID)
}
