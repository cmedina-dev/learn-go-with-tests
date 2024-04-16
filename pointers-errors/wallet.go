package pointers_errors

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Balance retrieves the current balance of the Wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit increases the current balance of the Wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw decreases the current balance of the Wallet while checking for insufficient funds.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// String formats the output for Bitcoin.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
