package state

import "fmt"

// State interface
type NCState interface {
	PressButton(*LightBulbNC) // Pass the light bulb as a parameter
}

// OnState represents the light being on
type OnStateNC struct{}

func (s *OnStateNC) PressButton(lightBulb *LightBulbNC) {
	fmt.Println("Turning the light off...")
	lightBulb.SetState(lightBulb.offState) // LightBulb manages its own state transition
}

// OffState represents the light being off
type OffStateNC struct{}

func (s *OffStateNC) PressButton(lightBulb *LightBulbNC) {
	fmt.Println("Turning the light on...")
	lightBulb.SetState(lightBulb.onState)
}

// LightBulb is the context class that switches between states
type LightBulbNC struct {
	onState  NCState
	offState NCState
	state    NCState
}

func NewLightBulbNC() *LightBulbNC {
	bulb := &LightBulbNC{}
	bulb.onState = &OnStateNC{}
	bulb.offState = &OffStateNC{}
	bulb.state = bulb.offState // Initial state: off
	return bulb
}

func (l *LightBulbNC) PressButton() {
	l.state.PressButton(l) // Pass the context (LightBulb) to the state
}

func (l *LightBulbNC) SetState(s NCState) {
	l.state = s
}
