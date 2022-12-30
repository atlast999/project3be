package test

import (
	"testing"
	"time"

	"github.com/atlast999/project3be/helper"
	"github.com/atlast999/project3be/token"
	"github.com/stretchr/testify/require"
)

func testTokenMaker(t *testing.T, tokenMaker token.TokenMaker) {
	userId := helper.RandInt(1, 100)
	username := helper.RandString(10)
	token, err := tokenMaker.CreateToken(userId, username, time.Hour)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	payload, err := tokenMaker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, userId, payload.UserID)
	require.Equal(t, username, payload.Username)
}

func TestJWTTokenMaker(t *testing.T) {
	secretKey := helper.RandString(32)
	tokenMaker, err := token.NewJWTMaker(secretKey)
	require.NoError(t, err)
	require.NotEmpty(t, tokenMaker)
	testTokenMaker(t, tokenMaker)
}

func TestPasetoTokenMaker(t *testing.T) {
	secretKey := helper.RandString(32)
	tokenMaker, err := token.NewPasetoMaker(secretKey)
	require.NoError(t, err)
	require.NotEmpty(t, tokenMaker)
	testTokenMaker(t, tokenMaker)
}
