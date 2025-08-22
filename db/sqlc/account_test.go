// An unit test for create account function
package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "Alamin",
		Balance:  20,
		Currency: "USD",
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

}
