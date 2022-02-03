package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.RWMutex // a binary semaphore guarding balance
	balance int
)

var endDepo = make(chan bool)
var endW = make(chan bool)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func deposit(amount int) {
	balance += amount
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()

	return balance
}

func doDepo(c int) {
	for i := 1; i <= c; i++ {
		Deposit(i * 100)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Deposited: %d\n", i*100)
	}
	endDepo <- true
}

func doWith(c int) {
	for i := 1; i <= c; i++ {
		s := Withdraw(i * 100)
		if s {
			fmt.Printf("Withdrew: %d\n", i*100)
		} else {
			fmt.Printf("Failed to withdraw: %d\n", i*100)
		}
		time.Sleep(100 * time.Millisecond)
	}
	endW <- true
}

func interfere(speed time.Duration) {
	for {
		select {
		case <-endDepo:
			<-endW
			return
		case <-endW:
			<-endDepo
			return
		default:
			time.Sleep(speed * time.Millisecond)
			fmt.Printf("Checking balance: %d\n", Balance())
		}
	}
}

func main() {
	go doDepo(20)
	go doWith(20)
	go interfere(300)

	select {
	case <-endDepo:
		<-endW
		fmt.Println("Gud baj 1")
		return
	case <-endW:
		<-endDepo
		fmt.Println("Gud baj 2")
		return
	}
}
