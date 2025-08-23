// An unit test for create account function
package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alaminpu1007/simplebank/util"
	"github.com/stretchr/testify/require"
)

// This method will skip from test, cause it's name is not start with 'Test'
func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testingQueries.CreateAccount(context.Background(), arg)

	// to test the result, install testify packages
	// https://github.com/stretchr/testify
	require.NoError(t, err)
	require.NotEmpty(t, account)

	// check equality after insert successfully
	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

/*
This method will test the creation procedure of account.
NOTE: Insert a dummy data on DB
*/
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

/*
This method will test the get procedure of account from db
*/
func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testingQueries.GeAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account1.Owner, account2.Owner)
	require.NotEmpty(t, account1.Balance, account2.Balance)
	require.NotEmpty(t, account1.Currency, account2.Currency)
	require.NotEmpty(t, account1.ID, account2.ID)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

/*
This method will test the update procedure of account
*/
func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testingQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, account2.ID, account2.ID)
	require.NotEmpty(t, account1.Owner, account2.Owner)
	require.NotEmpty(t, account2.Balance, arg.Balance)
	require.NotEmpty(t, account1.Currency, account2.Currency)
}

/*
This method will test the delete procedure of account
*/
func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testingQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	// now get item query
	accoun2, err := testingQueries.GeAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accoun2)
}

/*
This method will test the get lists of accounts
*/
func TestGetLists(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testingQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	// check exact 5 len
	require.Len(t, accounts, 5)

	// check any of them is empty or not
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
