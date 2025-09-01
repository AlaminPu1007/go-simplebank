package db

import (
	"context"
	"testing"
	"time"

	"github.com/alaminpu1007/simplebank/util"
	"github.com/stretchr/testify/require"
)

// This method will skip from test, cause it's name is not start with 'Test'
func createRandomUser(t *testing.T) User {
	hashPassword, err := util.HashedPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testingQueries.CreateUser(context.Background(), arg)

	// to test the result, install testify packages
	// https://github.com/stretchr/testify
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// check equality after insert successfully
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)
	require.Equal(t, user.FullName, arg.FullName)
	require.Equal(t, user.Email, arg.Email)

	require.True(t, user.PasswordChangeAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

/*
This method will test the creation procedure of user.
NOTE: Insert a dummy data on DB
*/
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

/*
This method will test the get procedure of user from db
*/
func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testingQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangeAt, user2.PasswordChangeAt, time.Second)
}
