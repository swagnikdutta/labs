package wallet

import (
	"errors"
	"testing"
)

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("Wasn't expecting any error but got one: %q", got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		// t.Fatal will stop the test if it is called.
		// If we don't stop, the following checks will cause panic since got is nil
		t.Fatal("wanted an error but didn't get one")
	}
	if !errors.Is(want, got) {
		t.Errorf("got %q, wanted %q", got, want) // not sure why type error works with %q
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		// TIL: If the type implements Stringer interface, we can use the %s format string with it
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(15))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})
	t.Run("no balance", func(t *testing.T) {
		startingBalance := Bitcoin(50)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
