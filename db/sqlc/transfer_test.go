package db

import (
	"context"
	"testing"
	"time"

	"github.com/alaminpu1007/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transers, err := testingQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transers)
	require.Equal(t, arg.FromAccountID, transers.FromAccountID)
	require.Equal(t, arg.ToAccountID, transers.ToAccountID)
	require.Equal(t, arg.Amount, transers.Amount)

	require.NotZero(t, transers.ID)
	require.NotZero(t, transers.CreatedAt)

	return transers
}

/*
This method will test the creation procedure of transfers.
NOTE: Insert a dummy data on DB
*/
func TestCreateTransers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	createRandomTransfer(t, account1, account2)
}

// This method will run test of get transer
func TestGetTranser(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	transfer2, err := testingQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

/*
This method will test the get lists of transfers
*/
func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transers, err := testingQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	// check exact 5 len
	require.Len(t, transers, 5)

	// check any of them is empty or not
	for _, transfer := range transers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
