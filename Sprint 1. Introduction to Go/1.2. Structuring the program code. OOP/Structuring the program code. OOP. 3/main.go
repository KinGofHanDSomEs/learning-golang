package main

import "fmt"

type Account struct {
	balance float64
	owner   string
}

func NewAccount(owner string) *Account {
	return &Account{balance: 0, owner: owner}
}

func (a *Account) SetBalance(balance float64) error {
	if balance < 0 {
		return fmt.Errorf("set balance on negative")
	}
	a.balance = balance
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(money float64) error {
	if money < 0 {
		return fmt.Errorf("deposit negative money")
	}
	a.balance += money
	return nil
}

func (a *Account) Withdraw(money float64) error {
	if money < 0 {
		return fmt.Errorf("negative withdraw")
	}
	if a.balance-money < 0 {
		return fmt.Errorf("withdrawal of a large amount of money than in the account")
	}
	a.balance -= money
	return nil
}
