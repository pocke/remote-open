package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	var port int
	var host string
	flag.IntVar(&port, "port", 2489, "TCP port number")
	flag.StringVar(&host, "host", "localhost", "Remote hostname")
	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, flag.Arg(0))
	conn.Write([]byte{'\000'})
	body, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	if len(body) != 0 {
		panic(string(body))
	}
}
