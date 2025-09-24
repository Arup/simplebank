package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}
	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		//require.Equal(t, amount, transfer.amount)
		// require.NotZero(t, transfer.ID)
		// require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check entries
		from_entry := result.FromEntry
		// result.NotEmpty(t, from_entry)
		// require.Equal(t, account1.ID, from_entry.AccountID)
		// require.Equal(t, -amount, from_entry.Amount)
		// require.NotZero(t, from_entry.ID)
		// require.NotZero(t, from_entry.CreatedAt)

		_, err = store.GetEntry(context.Background(), from_entry.ID)
		require.NoError(t, err)

		//check entries
		//to_entry := result.FromEntry
		// result.NotEmpty(t, to_entry)
		// require.Equal(t, account2.ID, to_entry.AccountID)
		// require.Equal(t, amount, to_entry.Amount)
		// require.NotZero(t, to_entry.ID)
		// require.NotZero(t, to_entry.CreatedAt)

		//TODO check accounts balance
	}
}
