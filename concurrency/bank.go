// Package bank provides a concurrency-safe bank with one account.
package main

import (
	"fmt"
	"time"
)

type withdrawReq struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)            // send amount to deposit
var balances = make(chan int)            // receive balance
var withdrawals = make(chan withdrawReq) // send amount to withdraw
var endDepo = make(chan bool)
var endW = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Withdraw(w withdrawReq) bool {
	withdrawals <- w

	return <-w.success
}
func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Printf("Added balance: %d \n", amount)
		case w := <-withdrawals:
			if balance < w.amount {
				w.success <- false
			} else {
				balance -= w.amount
				w.success <- true
			}
		case balances <- balance:
			fmt.Printf("Balance: %d\n", balance)
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func doDepo(count int, stop chan<- bool) {
	for i := 1; i <= count; i++ {
		Deposit(i * 100)
		time.Sleep(100 * time.Millisecond)
	}

	stop <- true
}

func doWith(count int, stop chan<- bool) {
	for i := 1; i <= count; i++ {
		success := Withdraw(withdrawReq{
			amount:  i * 100,
			success: make(chan bool),
		})
		fmt.Printf("Transaction success: %t - %d\n", success, i*100)
		time.Sleep(100 * time.Millisecond)
	}

	stop <- true
}

func main() {
	go doDepo(20, endDepo)
	go doWith(20, endW)

	select {
	case <-endDepo:
		<-endW
		Balance()
		time.Sleep(100 * time.Millisecond)
		return
	case <-endW:
		<-endDepo
		Balance()
		time.Sleep(100 * time.Millisecond)
		return
	}
}
