package main

type Wallet struct {
	balance int
}

// 不同之处在于，接收者类型是 *Wallet 而不是 Wallet，你可以将其解读为「指向 wallet 的指针」。
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}

// func (w Wallet) Deposit(amount int) {
// 	fmt.Println("address of balance in Deposit is", &w.balance)
// 	w.balance += amount
// }

// func (w Wallet) Balance() int {
// 	return w.balance
// }
