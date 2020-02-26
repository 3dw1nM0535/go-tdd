package pointers

import "testing"

func TestPointer(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("Widthraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Widthraw(Bitcoin(19))

		got := wallet.Balance()
		want := Bitcoin(1)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
