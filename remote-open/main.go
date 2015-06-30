package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/mitchellh/go-homedir"
	"github.com/monochromegane/conflag"
)

func main() {
	var port int
	var host string
	flag.IntVar(&port, "port", 2489, "TCP port number")
	flag.StringVar(&host, "host", "localhost", "Remote hostname")

	confPath, err := homedir.Expand("~/.config/remote-open.toml")
	if err != nil {
		panic(err)
	}
	if args, err := conflag.ArgsFrom(confPath); err == nil {
		flag.CommandLine.Parse(args)
	}

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
