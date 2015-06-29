package main

import (
	"bufio"
	"log"
	"net"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	l, err := net.Listen("tcp", "localhost:2489")
	if err != nil {
		panic(err)
	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	log.Printf("Request from %s", conn.RemoteAddr())
	line, err := bufio.NewReader(conn).ReadString('\000')
	body := line[:len(line)-1]

	if err != nil {
		log.Println(err)
		conn.Write([]byte(err.Error()))
		return
	}

	err = open.Run(string(body))
	if err != nil {
		log.Println(err)
		conn.Write([]byte(err.Error()))
		return
	}
}
