package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// RUN A TEST METHOD THAT WILL TEST THE TRANSFER TRANSACTION
func TestTranserTx(t *testing.T) {

	// create a new store
	store := NewStore(testDB)

	// we will send money from account1 to account2
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	fmt.Println(">> before:", account1.Balance, account2.Balance)

	/*
		writing database transaction is something we must always be very careful with.
		It can be easy to write, but can also easily become a nightmare if we don’t handle the concurrency carefully.
		So the best way to make sure that our transaction works well is to run it with several concurrent go routines.
	*/
	n := 10

	amount := int64(10)

	// Run n concurrent transfer transaction

	// create or initialized a channel, to hold or get back the err or result from go routine
	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		// log to detect deadlock after convenrtion get account query using "FOR UPDATE"
		txName := fmt.Sprintf("tx %d", i+1)

		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)

			result, err := store.TrnasterTx(ctx, TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)

	// check results
	for i := 0; i < n; i++ {
		err := <-errs

		// if any err is present
		require.NoError(t, err)

		result := <-results

		// result should not be empty
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer

		// check coupls of conditions
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, account1.ID)
		require.Equal(t, transfer.ToAccountID, account2.ID)
		require.Equal(t, transfer.Amount, amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		// now check the transfer is really created or inserted into data-base
		_, err = store.GetTransfer(context.Background(), transfer.ID)

		// if any err is present
		require.NoError(t, err)

		// Check account entries
		fromEntry := result.FromEntry

		// check coupls of conditions
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		// now check the transfer is really created or inserted into data-base
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		// CHECK ENTRY 2
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)

		// if any err is present
		require.NoError(t, err)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// check balances
		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance

		// write test for diff
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		existed[k] = true

	}

	// ref: https://dev.to/techschoolguru/db-transaction-lock-how-to-handle-deadlock-22o8

	// check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

// CREATE A DEAD LOCK SENARIO
func TestTrasferDeadLock(t *testing.T) {

	// create a new store
	store := NewStore(testDB)

	// we will send money from account1 to account2
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	/*
		writing database transaction is something we must always be very careful with.
		It can be easy to write, but can also easily become a nightmare if we don’t handle the concurrency carefully.
		So the best way to make sure that our transaction works well is to run it with several concurrent go routines.
	*/
	n := 10

	amount := int64(10)

	// Run n concurrent transfer transaction

	// create or initialized a channel, to hold or get back the err or result from go routine
	errs := make(chan error)

	for i := 0; i < n; i++ {

		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			_, err := store.TrnasterTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	// check errors
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// ref: https://dev.to/techschoolguru/how-to-avoid-deadlock-in-db-transaction-queries-order-matter-oh7

	// check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}
