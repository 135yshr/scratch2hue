package scratch2hue

import (
	scratch "github.com/135yshr/scratchgo"
	"github.com/savaki/go.hue"
)

const (
	typeBroadcast = "broadcast"
	typeSensor    = "sensor-update"
)

type HueConnection struct {
	bridge *hue.Bridge
	state  hue.SetLightState
	id     string
}

func NewConnection(ipaddr string) *HueConnection {
	bridge := hue.NewBridge(ipaddr, "scratchdev")
	var state hue.SetLightState
	return &HueConnection{bridge: bridge, state: state}
}

func (self *HueConnection) Anction(msg *scratch.Message) error {
	action := self.create_action_type(msg.Type)
	return action(msg)
}

func (self *HueConnection) create_action_type(message_type string) func(msg *scratch.Message) error {
	switch message_type {
	case typeBroadcast:
		return func(msg *scratch.Message) error {
			switch msg.Variables["command"] {
			case "action":
				return self.broadcast_action()

			case "discotime":
				return self.broadcast_discotime()

			case "light_on":
				return self.broadcast_light_on()

			case "light_off":
				return self.broadcast_light_off()

			case "light_all_on":
				return self.broadcast_light_all_on()

			case "light_all_off":
				return self.broadcast_light_all_off()
			}
			return nil
		}
	case typeSensor:
		return func(msg *scratch.Message) error {
			val := msg.Variables
			for _, v := range msg.GetNames() {
				switch v {
				case "on":
					self.state.On = val["on"]
				case "brightness":
					self.state.Bri = val["brightness"]
				case "color":
					self.state.Hue = val["color"]
				case "id":
					self.id = val["id"]
				}
			}
			return nil
		}
	}
	return nil
}

func (self *HueConnection) broadcast_action() error {
	light, err := self.bridge.FindLightById(self.id)
	if err != nil {
		return err
	}

	_, err = light.SetState(self.state)
	if err != nil {
		return err
	}
	return nil
}

func (self *HueConnection) broadcast_discotime() error {
	light, err := self.bridge.FindLightById(self.id)
	if err != nil {
		return err
	}

	_, err = light.ColorLoop()
	if err != nil {
		return err
	}
	return nil
}

func (self *HueConnection) broadcast_light_on() error {
	light, err := self.bridge.FindLightById(self.id)
	if err != nil {
		return err
	}

	_, err = light.On()
	if err != nil {
		return err
	}
	return nil
}

func (self *HueConnection) broadcast_light_off() error {
	light, err := self.bridge.FindLightById(self.id)
	if err != nil {
		return err
	}

	_, err = light.Off()
	if err != nil {
		return err
	}
	return nil
}

func (self *HueConnection) broadcast_light_all_on() error {
	lights, err := self.bridge.GetAllLights()
	if err != nil {
		return err
	}

	for _, l := range lights {
		_, err := l.On()
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *HueConnection) broadcast_light_all_off() error {
	lights, err := self.bridge.GetAllLights()
	if err != nil {
		return err
	}

	for _, l := range lights {
		_, err := l.Off()
		if err != nil {
			return err
		}
	}
	return nil
}
