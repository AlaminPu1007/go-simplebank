package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

// return new Store object
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execute a generic database transaction.
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {

	//  start a new transaction
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

// TransferTxParams -> contains the inut parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxParams -> the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

// TransferTX -> Perform a money transfer from one account to the other
// It will creates a transfer record, add account entries, and upate accounts balance withing a single data transaction
func (store *Store) TransferTX(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// txName := ctx.Value(txKey)

		// log to detect deadlock after convenrtion get account query using "FOR UPDATE"
		// fmt.Println(txName, "create transfer")

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// log to detect deadlock after convenrtion get account query using "FOR UPDATE"
		// fmt.Println(txName, "create entry 1")

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// log to detect deadlock after convenrtion get account query using "FOR UPDATE"
		// fmt.Println(txName, "create entry 2")

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// log to detect deadlock after convenrtion get account query using "FOR UPDATE"
		// fmt.Println(txName, "Get account 1")

		// move money out of account1
		// account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountID) // converth GetAccount to GetAccountForUpdate to commit db transaction
		// if err != nil {
		// 	return err
		// }

		// log to detect deadlock after convenrtion get account query using "FOR UPDATE"
		// fmt.Println(txName, "update account 1")

		/*
			WE ALWAYS UPDATE SMALLER ACCOUNT FIRST, TO AVOID DEAD LOCK SITUATION.
		*/
		// if arg.FromAccountID < arg.ToAccountID {
		// 	result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		// 		ID:     arg.FromAccountID,
		// 		Amount: -arg.Amount,
		// 	})
		// 	if err != nil {
		// 		return err
		// 	}

		// log to detect deadlock after convenrtion get account query using "FOR UPDATE"
		// fmt.Println(txName, "get account 2")

		// move money into account2
		// account2, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
		// if err != nil {
		// 	return err
		// }

		// log to detect deadlock after convenrtion get account query using "FOR UPDATE"
		// fmt.Println(txName, "update account 2")

		// result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		// 	ID:     arg.ToAccountID,
		// 	Amount: arg.Amount,
		// })
		// if err != nil {
		// 	return err
		// }
		// } else {

		// 	result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		// 		ID:     arg.ToAccountID,
		// 		Amount: arg.Amount,
		// 	})
		// 	if err != nil {
		// 		return err
		// 	}

		// 	result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		// 		ID:     arg.FromAccountID,
		// 		Amount: -arg.Amount,
		// 	})
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		// WE ALWAYS UPDATE SMALLER ACCOUNT FIRST, TO AVOID DEAD LOCK SITUATION
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return err
	})

	return result, err
}

// TO REFACTOR CODE OR REMOVING DUPLICATE,
func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64) (account1 Account, account2 Account, err error) {

	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		// implecitly return : account1 Account, account2 Account, err error
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	// implecitly return : account1 Account, account2 Account, err error
	return
}
