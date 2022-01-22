package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func echo(c net.Conn, shout string, delay time.Duration, wg sync.WaitGroup) {
	defer wg.Done()

	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

const timeout = time.Second * 5

func handleConn(c net.Conn) {
	backoff := time.After(timeout)
	cc := make(chan string)
	wg := sync.WaitGroup{}
	input := bufio.NewScanner(c)

	defer func() {
		wg.Wait()
		c.Close()
	}()

	go func() {
		fmt.Println("Joined: ", c)
		for input.Scan() {
			fmt.Println("Input scan", input.Text())
			cc <- input.Text()
		}
		if input.Err() != nil {
			log.Print("Scan: ", input.Err())
		}
	}()

	select {
	case <-backoff:
		fmt.Printf("\rClosing connection: %v", c)
		wg.Done()
		c.Close()
	case <-cc:
		backoff = time.After(timeout)
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, wg)
	}
}
