package main

import "fmt"

type Account interface {
	Deposit(amount int)
	Balance() int
}

type SavingsAccount struct {
	balance int
}

func (s *SavingsAccount) Deposit(amount int) {
	s.balance = s.balance + amount
}

func (ss SavingsAccount) Balance() int {
	return ss.balance
}

type CreditAccount struct {
	balance     int
	Creditlimit int
}

func (c *CreditAccount) Deposit(amount int) {
	if c.balance+c.Creditlimit > 0 {
		c.balance = c.balance + amount
	}

}
func (cc CreditAccount) Balance() int {
	return cc.balance
}

func printBalance(a Account) {
	fmt.Println(a.Balance())
}

func main() {
	a := SavingsAccount{balance: 1000}
	b := CreditAccount{balance: 1000, Creditlimit: 10000}

	printBalance(&a)
	printBalance(&b)

	a.Deposit(500)
	b.Deposit(700)

	printBalance(&a)
	printBalance(&b)

	b.Deposit(12000)
	printBalance(&b)
	b.Deposit(-15000)
}
