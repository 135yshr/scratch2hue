package main

import (
	"flag"
	"fmt"
	scratch "github.com/135yshr/scratchgo"
	"github.com/savaki/go.hue"
	"os"
)

var ipaddr string
var state hue.SetLightState
var id string

func main() {

	flag.Parse()

	conn, bridge := initialize_routine()
	for {
		msg, err := conn.Recv()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*msg)
		action(msg, bridge)
	}

	os.Exit(0)
}

func init() {
	flag.StringVar(&ipaddr, "ip", "not found", "bridge ip address.")
}

func action(msg *scratch.Message, bridge *hue.Bridge) {
	switch msg.Type {
	case "broadcast":
		light, err := bridge.FindLightById(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch msg.Variables["command"] {
		case "action":
			results, err := light.SetState(state)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(results)

		case "discotime":
			results, err := light.ColorLoop()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(results)

		case "light_on":
			results, err := light.On()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(results)

		case "light_off":
			results, err := light.Off()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(results)

		case "light_all_on":
			lights, err := bridge.GetAllLights()
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, l := range lights {
				results, err := l.On()
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(results)
			}
		case "light_all_off":
			lights, err := bridge.GetAllLights()
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, l := range lights {
				results, err := l.Off()
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(results)
			}
			fmt.Println("success!")
		}
	case "sensor-update":
		val := msg.Variables
		for _, v := range msg.GetNames() {
			switch v {
			case "on":
				state.On = val["on"]
			case "brightness":
				state.Bri = val["brightness"]
			case "color":
				state.Hue = val["color"]
			case "id":
				id = val["id"]
			}
		}
	}
}

func errorToExit(message string, err error) {
	if err != nil {
		fmt.Println(message)
		fmt.Println(err)
		os.Exit(1)
	}
}

func errorToConsole(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func initialize_routine() (*scratch.RspConn, *hue.Bridge) {
	fmt.Println("now during initializing...")

	conn, err := scratch.NewDefaultConnect()
	errorToExit("connect to scratch.", err)

	bridge := hue.NewBridge(ipaddr, "newdeveloper")
	fmt.Println("registered new device =>", *bridge)

	return conn, bridge
}
