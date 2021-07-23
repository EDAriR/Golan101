package main

import (
	"errors"
	"fmt"
)

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

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("oh no")
	}

	w.balance -= amount
	return nil
}

// 不同之处在于，接收者类型是 *Wallet 而不是 Wallet，你可以将其解读为「指向 wallet 的指针」。
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// func (w Wallet) Deposit(amount int) {
// 	fmt.Println("address of balance in Deposit is", &w.balance)
// 	w.balance += amount
// }

// func (w Wallet) Balance() int {
// 	return w.balance
// }
