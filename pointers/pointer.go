package pointers

import (
	"fmt"

	"errors"
)

// Bitcoin : type bitcoin
type Bitcoin int

// Stringer : interface for printing to the StdOut
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet : type wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit : add funds to your account
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance : returns the account balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrorInsufficientBalance : error message for insufficient funds during withdrawal
var ErrorInsufficientBalance = errors.New("Oh no! Insufficient balance in your account to initiate widthrawal")

// Widthraw : widthraw cash from the account
func (w *Wallet) Widthraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrorInsufficientBalance
	}
	w.balance -= amount
	return nil
}
