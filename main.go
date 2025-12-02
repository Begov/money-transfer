package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(coins float64) {
	u.Balance += coins
	fmt.Printf("Счёт пользователя %s пополнен на %.2f\n", u.Name, coins)
}

func (u *User) Withdraw(coins float64) error {
	if u.Balance-coins < 0 {
		return errors.New("insufficient funds on balance")
	}
	u.Balance -= coins
	return nil
}

func main() {
	user1 := User{ID: "1", Name: "Иван", Balance: 400}
	user2 := User{ID: "2", Name: "Артем", Balance: 1268.04}

	user1.Deposit(200)

	if err := user1.Withdraw(700); err != nil {
		fmt.Printf("%s: %s\n", user1.Name, err)
	}

	if err := user2.Withdraw(700); err != nil {
		fmt.Printf("%s: %s\n", user2.Name, err)
	}

	fmt.Printf("%s: %.2f на балансе.\n", user1.Name, user1.Balance)
	fmt.Printf("%s: %.2f на балансе.\n", user2.Name, user2.Balance)
}
