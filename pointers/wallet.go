package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	amount Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient amount")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.amount += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.amount < amount {
		return ErrInsufficientFunds
	}
	w.amount -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.amount
}
