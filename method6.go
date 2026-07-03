package main

import "fmt"

type Deposit struct {
	Owner string

	Balance int
}

func (a *Deposit) Depo(amount int) {
	a.Balance = amount + a.Balance
}
func (w *Deposit) Withdraw(amount int) bool {
	if w.Balance >= amount {
		w.Balance = w.Balance - amount
		return true
	} else {
		fmt.Println("недостаточно средств")
		return false
	}
}

func (p Deposit) Print() {
	fmt.Println(p.Owner, ": ", p.Balance)
}
func main() {
	p := Deposit{Owner: "Артур", Balance: 1000}
	p.Print()
	p.Depo(300)
	p.Print()
	p.Withdraw(1000)
	p.Print()
	p.Withdraw(2000)
	p.Print()
}
