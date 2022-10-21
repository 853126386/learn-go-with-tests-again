package pointer

import (
	"errors"
	"fmt"
)
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")


type Stringer interface {
	String() string
}

type Btc int

type Wallet struct {
	blance Btc
}

func (w *Wallet) Deposit(amount Btc)  {
	w.blance +=amount
}

func (w *Wallet) Withdraw(amount Btc) error  {
	if w.blance < amount {
		return InsufficientFundsError
	}
	w.blance -=amount
	return nil
}

func (w *Wallet) Blance()Btc  {
	return w.blance
}

func (b Btc) String()string  {
	return fmt.Sprintf("%d Btc",b)
}
