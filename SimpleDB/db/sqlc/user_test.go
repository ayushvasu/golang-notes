package db

import (
	"context"
	"simplebank/db/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	var username = util.RandomOwner()
	var hashedPassword = "secret"
	var fullName = util.RandomOwner()
	var email = util.RandomEmail()

	arg := CreateUserParams{
		Username:       username,
		HashedPassword: hashedPassword,
		FullName:       fullName,
		Email:          email,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, username, user.Username)
	require.Equal(t, hashedPassword, user.HashedPassword)
	require.Equal(t, fullName, user.FullName)
	require.Equal(t, email, user.Email)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
