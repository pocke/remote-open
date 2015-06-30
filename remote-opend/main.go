package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mitchellh/go-homedir"
	"github.com/monochromegane/conflag"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 2489, "TCP port number")

	confPath, err := homedir.Expand("~/.config/remote-opend.toml")
	if err != nil {
		panic(err)
	}
	if args, err := conflag.ArgsFrom(confPath); err == nil {
		flag.CommandLine.Parse(args)
	}

	flag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
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

	if err != nil {
		log.Println(err)
		conn.Write([]byte(err.Error()))
		return
	}
	body := line[:len(line)-1]

	err = open.Run(string(body))
	if err != nil {
		log.Println(err)
		conn.Write([]byte(err.Error()))
		return
	}
}
