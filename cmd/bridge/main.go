package main

import (
	"fmt"
	"go-design-patterns/pkg/bridge"
)

func main() {
	// Use TV with the remote control
	tv := &bridge.TV{}
	remote := &bridge.RemoteControl{Device: tv}

	remote.TurnOn()
	remote.VolumeUp()
	remote.VolumeUp()
	remote.VolumeDown()
	remote.TurnOff()

	fmt.Println()

	// Use Radio with the advanced remote control
	radio := &bridge.Radio{}
	advancedRemote := &bridge.AdvancedRemoteControl{bridge.RemoteControl{Device: radio}}

	advancedRemote.TurnOn()
	advancedRemote.VolumeUp()
	advancedRemote.Mute()
	advancedRemote.TurnOff()
}
