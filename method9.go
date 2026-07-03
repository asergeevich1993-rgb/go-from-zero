package main

import "fmt"

type Wallet struct {
	Balance int
}

func (b *Wallet) Add(amount int) {
	b.Balance = b.Balance + amount

}
func (s *Wallet) Spend(amount int) bool {
	if s.Balance >= amount {
		s.Balance = s.Balance - amount
		return true
	} else {
		return false
	}

}
func (p Wallet) Print() {
	fmt.Println("Баланс :", p.Balance)

}

func main() {

	wallet := Wallet{Balance: 1000}

	wallet.Print()

	wallet.Add(500)

	wallet.Print()

	if wallet.Spend(700) {
		fmt.Println("Покупка успешна", wallet.Balance)

	}
	wallet.Print()
	if wallet.Spend(3000) == false {
		fmt.Println("Недостаточно средств на балансе: ", wallet.Balance)
	}

}
