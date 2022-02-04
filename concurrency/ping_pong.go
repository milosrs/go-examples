package main

import (
	"fmt"
	"sync"
	"time"
)

const timeout = 1 * time.Second

var (
	m     sync.RWMutex
	pongs = 0
)

func ping(req chan<- struct{}, done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			req <- struct{}{}
		}
	}
}

func pong(resp <-chan struct{}) {
	for _ = range resp {
		m.Lock()
		pongs++
		m.Unlock()
	}
}

func main() {
	pingPong := make(chan struct{})
	done := make(chan bool)
	timer := time.After(timeout)

	go ping(pingPong, done)
	go pong(pingPong)

	select {
	case <-timer:
		m.RLock()
		fmt.Printf("Communicated %d times!", pongs)
		m.RUnlock()
		done <- true
		close(pingPong)
		return
	}
}
