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

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

type PaymentSystem struct {
	Users        map[string]User
	Transactions []Transaction
}

func (ps *PaymentSystem) AddUser(u User) {
	ps.Users[u.ID] = u
}

func (ps *PaymentSystem) AddTransaction(t Transaction) {
	ps.Transactions = append(ps.Transactions, t)
}

func (ps *PaymentSystem) ProcessingTransactions() error {
	for _, t := range ps.Transactions {
		fromUser, fromExist := ps.Users[t.FromID]
		toUser, toExist := ps.Users[t.ToID]

		if !fromExist {
			return fmt.Errorf("user with ID %s not found", t.FromID)
		}

		if !toExist {
			return fmt.Errorf("user with ID %s not found", t.ToID)
		}

		if err := fromUser.Withdraw(t.Amount); err != nil {
			return fmt.Errorf("Error: %v", err)
		}

		toUser.Deposit(t.Amount)

		// ps.Users[t.FromID] = fromUser
		// ps.Users[t.ToID] = toUser
	}

	ps.Transactions = nil
	return nil
}

func main() {

	ps := &PaymentSystem{
		Users:        make(map[string]User),
		Transactions: []Transaction{},
	}

	user1 := &User{ID: "1", Name: "Иван", Balance: 400}
	user2 := &User{ID: "2", Name: "Артем", Balance: 1268.04}

	ps.AddUser(*user1)
	ps.AddUser(*user2)

	ps.AddTransaction(Transaction{"1", "2", 200})
	ps.AddTransaction(Transaction{"2", "1", 50})

	if err := ps.ProcessingTransactions(); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

}
