package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserAccount struct {
	sync.Mutex
	Balance int
	Name    string
}

func (account *UserAccount) Lock() {
	account.Mutex.Lock()
}

func (account *UserAccount) Unlock() {
	account.Mutex.Unlock()
}

func (account *UserAccount) Change(amount int) {
	account.Balance += amount
}

func Transfer(user1 *UserAccount, user2 *UserAccount, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	defer user1.Unlock()
	defer user2.Unlock()
}

// situasi deadlock membuat balance tidak sesuai
func TestTransferDeadlock(t *testing.T) {
	budi := UserAccount{Balance: 100, Name: "Budi"}
	kerel := UserAccount{Balance: 1000, Name: "Kerel"}

	go Transfer(&budi, &kerel, 50)
	go Transfer(&kerel, &budi, 400)

	time.Sleep(5 * time.Second)
	fmt.Println("Budi Balance:", budi.Balance)
	fmt.Println("Kerel Balance:", kerel.Balance)
}
