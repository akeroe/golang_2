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

	ps := &PaymentSystem{
		Users: make(map[string]*User),
	}
	ps.AddUser(user1)
	fmt.Println(ps.Users["1"].Balance)
	ps.AddUser(user2)
	fmt.Println(ps.Users["2"].Balance)

	t1 := Transaction{
		FromID: "1",
		ToID:   "2",
		Amount: 200,
	}
	t2 := Transaction{
		FromID: "2",
		ToID:   "1",
		Amount: 50,
	}
	t3 := Transaction{
		FromID: "999",
		ToID:   "2",
		Amount: 100,
	}




	
	ps.AddTransaction(t1)
	ps.AddTransaction(t2)
	ps.AddTransaction(t3)
	fmt.Println("Транзакций в очереди:", len(ps.TransactionQueue))

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Transfer succesful")
	// }
	// err = ps.ProcessingTransaction(t2)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Transfer succesful")
	// }

	for _, t := range ps.TransactionQueue {
		err = ps.ProcessingTransaction(t)
		if err != nil {
			fmt.Println(err) // получили ошибку из withdraw
		} else {
			fmt.Println("Transfer successfull")
		}

	}
	fmt.Println("Баланс Max:", ps.Users["1"].Balance)
	fmt.Println("Баланс Alex:", ps.Users["2"].Balance)

}

type User struct {
	ID      string
	Name    string
	Balance float64
}

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

type PaymentSystem struct {
	Users            map[string]*User // хранение пользователей по ID. String - ключ (userid), *User - указатель на юзера
	TransactionQueue []Transaction    // очередь денежных переводов
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

func (ps *PaymentSystem) AddUser(user *User) {
	ps.Users[user.ID] = user
}
func (ps *PaymentSystem) AddTransaction(t Transaction) {
	ps.TransactionQueue = append(ps.TransactionQueue, t)
}

func (ps *PaymentSystem) ProcessingTransaction(t Transaction) error {
	sender, ok := ps.Users[t.FromID] // true or false существует ли ключ
	if !ok {
		return errors.New("sender not found")
	}
	receiver, ok := ps.Users[t.ToID]
	if !ok {
		return errors.New("receiver not found")
	}
	err := sender.Withdraw(t.Amount)
	if err != nil {
		return err
	}
	receiver.Deposit(t.Amount)

	return nil

}
