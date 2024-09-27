package bridge

import "fmt"

type Device interface {
	On()
	Off()
	SetVolume(volume int)
	GetVolume() int
}

type TV struct {
	Volume int
}

func (t *TV) On() {
	fmt.Println("TV is now ON")
}

func (t *TV) SetVolume(volume int) {
	t.Volume = volume
	fmt.Printf("TV volume set to %d\n", volume)
}

func (t *TV) Off() {
	fmt.Println("TV is now OFF")
}

func (t *TV) GetVolume() int {
	return t.Volume
}

// radio is another concrete implementor

type Radio struct {
	Volume int
}

func (r *Radio) On() {
	fmt.Println("Radio is now ON")
}

func (r *Radio) Off() {
	fmt.Println("Radio is now OFF")
}

func (r *Radio) SetVolume(volume int) {
	r.Volume = volume
	fmt.Printf("Radio volume set to %d\n", volume)
}

func (r *Radio) GetVolume() int {
	return r.Volume
}

// RemoteControl is the abstraction that delegates the operation to the device
type RemoteControl struct {
	Device Device
}

func (r *RemoteControl) TurnOn() {
	r.Device.On()
}

func (r *RemoteControl) TurnOff() {
	r.Device.Off()
}

func (r *RemoteControl) VolumeUp() {
	volume := r.Device.GetVolume()
	r.Device.SetVolume(volume + 1)
}

func (r *RemoteControl) VolumeDown() {
	volume := r.Device.GetVolume()
	r.Device.SetVolume(volume - 1)
}

// AdvancedRemoteControl is the refined abstraction
type AdvancedRemoteControl struct {
	RemoteControl // Embedding the basic RemoteControl
}

func (r *AdvancedRemoteControl) Mute() {
	r.Device.SetVolume(0)
	fmt.Println("Device is muted")
}
