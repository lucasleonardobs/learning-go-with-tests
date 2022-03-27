package wallet

import (
	"errors"
	"fmt"
)

var ErrorInsufficientBalance = errors.New("cannot withdraw: insufficient balance")

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return ErrorInsufficientBalance
	}

	w.balance -= value
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
