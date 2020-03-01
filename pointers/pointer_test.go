package pointers

import "testing"

func TestPointer(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but expected one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but expected none")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Widthraw sufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Widthraw(10)

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Widthraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Widthraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrorInsufficientBalance)

	})
}
