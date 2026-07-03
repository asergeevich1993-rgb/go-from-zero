package main

import "fmt"

type Account struct {
	Owner string

	Balance int
}

func (d *Account) Deposit(amount int) {
	d.Balance = d.Balance + amount
}
func (w *Account) Withdraw(amount int) bool {
	if w.Balance >= amount {
		w.Balance = w.Balance - amount
		return true
	} else {
		fmt.Println("недостаточно средств у,", w.Owner)
		return false
	}
}

func (pr Account) Print() {
	fmt.Println(pr.Owner, ":", pr.Balance)

}

func main() {
	accounts := []Account{
		{Owner: "Артур", Balance: 1000},
		{Owner: "Мария", Balance: 500},
		{Owner: "Иван", Balance: 2000}}

	for _, acc := range accounts {
		acc.Print()
	}
	accounts[0].Deposit(500)
	if accounts[1].Withdraw(1000) {
		fmt.Println("Снятие успешно")
	} else {
		fmt.Println("Снятие неуспешно")
	}
	if accounts[2].Withdraw(500) {
		fmt.Println("Снятие успешно")
	} else {
		fmt.Println("снятие неуспешно")
	}
	for _, acc := range accounts {
		acc.Print()
	}

}
