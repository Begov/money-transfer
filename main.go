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
	fmt.Println(user1.Balance)

	err := user1.Withdraw(700)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(user1.Balance)
	}

	user2.Deposit(200)
	fmt.Println(user2.Balance)

	err = user2.Withdraw(500)

	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println(user2.Balance)
	}

}
