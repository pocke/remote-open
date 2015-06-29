package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2489")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, os.Args[1])
	conn.Write([]byte{'\000'})
	body, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	if len(body) != 0 {
		panic(string(body))
	}
}
