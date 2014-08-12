package main

import (
	"flag"
	"fmt"
	scratch "github.com/135yshr/scratchgo"
	"os"
	s2hue "scratch2hue"
)

var ipaddr string

func main() {

	flag.Parse()

	if ipaddr == "" {
		fmt.Println("set ipaddrss")
		os.Exit(1)
	}

	conn, err := scratch.NewDefaultConnect()
	errorToExit("connect to scratch.", err)

	hue := s2hue.NewConnection(ipaddr)
	for {
		msg, err := conn.Recv()
		if err != nil {
			fmt.Println(err)
			continue
		}
		hue.Anction(msg)
	}

	os.Exit(0)
}

func init() {
	flag.StringVar(&ipaddr, "ip", "", "bridge ip address.")
}

func errorToExit(message string, err error) {
	if err != nil {
		fmt.Println(message)
		fmt.Println(err)
		os.Exit(1)
	}
}
