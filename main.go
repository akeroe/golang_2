package main

import (
	"errors"
	"fmt"
)

func main() {
	user1 := &User{
		Name:    "Max",
		ID:      "1",
		Balance: 1000,
	}

	user1.Deposit(200)
	fmt.Println("После пополнения:", user1.Name, user1.Balance)

	err := user1.Withdraw(1000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("После снятия:", user1.Name, user1.Balance)
	}

	user2 := &User{
		Name:    "Alex",
		ID:      "2",
		Balance: 500,
	}
	user2.Deposit(300)
	fmt.Println("После пополнения:", user2.Name, user2.Balance)

	err = user2.Withdraw(200)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("После снятия:", user2.Name, user2.Balance)
	}

}

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(dodep float64) { //dodep - сумма пополнения
	u.Balance += dodep
}

func (u *User) Withdraw(withdr float64) error { //withdr сколько снимает
	if withdr <= u.Balance {
		u.Balance -= withdr
		return nil

	} else {
		return errors.New("Error: insufficient funds")

	}
}
