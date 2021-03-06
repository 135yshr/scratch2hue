package main

import (
	"flag"
	"fmt"
	s2hue "github.com/135yshr/scratch2hue"
	scratch "github.com/135yshr/scratchgo"
	"os"
)

var ipaddr string

func main() {

	flag.Parse()

	if ipaddr == "" {
		fmt.Println("set ipaddrss")
		os.Exit(1)
	}

	conn, err := scratch.NewDefaultConnect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hue := s2hue.NewConnection(ipaddr)
	for {
		msg, err := conn.Recv()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if err = hue.Action(msg); err != nil {
			fmt.Println(err)
		}
	}

	os.Exit(0)
}

func init() {
	flag.StringVar(&ipaddr, "ip", "", "bridge ip address.")
}
