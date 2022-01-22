package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var isServer = flag.Bool("server", false, "server = true starts listening for messages - false behaves as a client")

const loc = "localhost:8600"
const protocol = "tcp"

var myCon string

func main() {
	flag.Parse()

	if *isServer {
		done := make(chan struct{})
		startTCPServer(loc, protocol, done)
		fmt.Println("Server started")
		<-done
	} else {
		con, err := dialChatServer(loc, protocol)
		written := make(chan struct{})
		if err != nil {
			log.Fatal(err)
		}
		defer con.Close()

		for {
			go msgReceiveHandler(con)
			msgSendHandler(con, written)
			<-written
		}
	}
}

type conns map[string]net.Conn

const exit = "exit"

func startTCPServer(location, connection string, done chan<- struct{}) {
	conList := make(conns, 0)
	listener, err := net.Listen(connection, location)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		for _, c := range conList {
			if err := c.Close(); err != nil {
				fmt.Fprintf(os.Stdout, "Error closing network: %v", err)
			}
		}
	}()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
			}
			conList[conn.RemoteAddr().String()] = conn
			go handleConn(conn, &conList)
		}
	}()

	go func() {
		var scan string
		fmt.Scan(scan)

		if scan == io.EOF.Error() {
			done <- struct{}{}
		}
	}()
}

func closeConn(c net.Conn, conList *conns) {
	c.Close()
	delete(*conList, c.LocalAddr().Network())
	fmt.Printf("Deleted %s\n", c.RemoteAddr().String())
}

func handleConn(conn net.Conn, conList *conns) {
	fmt.Println("New connection:", conn)
	bufreader := bufio.NewScanner(conn)

	for bufreader.Scan() {
		msg := bufreader.Text()
		fmt.Println(conn.RemoteAddr().String()+": ", msg)
		if msg == "exit" {
			closeConn(conn, conList)
		}

		for _, v := range *conList {
			if v.RemoteAddr().String() != conn.RemoteAddr().String() {
				fmt.Fprintf(v, "%s says: %s\n", conn.RemoteAddr().String(), msg)
			}
		}
	}
}

func dialChatServer(location, connection string) (net.Conn, error) {
	a, err := net.ResolveTCPAddr(connection, location)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP(connection, nil, a)
	if err != nil {
		log.Fatal(err)
	}

	return conn, err
}

func msgSendHandler(server net.Conn, written chan<- struct{}) {
	_, err := io.Copy(server, os.Stdin)
	if err != nil {
		fmt.Errorf("Failed sending message to server {%s}. Reason %v", server.LocalAddr().Network(), err)
	}
	written <- struct{}{}
}

func msgReceiveHandler(server net.Conn) {
	_, err := io.Copy(os.Stdout, server)
	if err != nil {
		fmt.Errorf("Failed sending message to server {%s}. Reason %v", server.LocalAddr().Network(), err)
	}
}
