package test

import (
	"testing"

	"github.com/atlast999/project3be/helper"
	"github.com/stretchr/testify/require"
)

func TestGeneratePassword(t *testing.T) {
	raw := helper.RandString(6)
	hashed, err := helper.GeneratePassword(raw)
	require.NoError(t, err)
	require.NotEmpty(t, hashed)
}

func TestCheckPasswrod(t *testing.T) {
	raw := helper.RandString(6)
	hashed, err := helper.GeneratePassword(raw)
	require.NoError(t, err)
	require.NotEmpty(t, hashed)
	err = helper.CheckPassword(raw, hashed)
	require.NoError(t, err)
	wrong := helper.RandString(7)
	err = helper.CheckPassword(wrong, hashed)
	require.Error(t, err)
}
