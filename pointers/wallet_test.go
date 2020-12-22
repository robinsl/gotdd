package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Can Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		want := Bitcoin(20)
		assertBalance(t, wallet, want)
	})

	t.Run("Can Withdraw", func(t *testing.T) {
		wallet := Wallet{amount: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw too much", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		got := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, got, ErrInsufficientFunds)
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Application should have fire an error since it is invalid")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("Application thrown an unexpected error")
	}
}

