package pointers

import (
	// "errors"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}

		err := wallet.Withdraw(Bitcoin(25))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(25))
	})

	t.Run("Withdraw insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(100)

		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(200))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error)  {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didnt want one.")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didnt get one.")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
