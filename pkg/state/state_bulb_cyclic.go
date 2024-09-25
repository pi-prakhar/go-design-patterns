package state

import "fmt"

// State interface defines the action that can change depending on the state
type CState interface {
	PressButton() // Action that will toggle the light
}

// OnState represents the light bulb being on
type OnStateC struct {
	lightBulb *LightBulbC
}

func (s *OnStateC) PressButton() {
	fmt.Println("Turning the light off...")
	s.lightBulb.SetState(s.lightBulb.offState) // Change to OffState
}

// OffState represents the light bulb being off
type OffStateC struct {
	lightBulb *LightBulbC
}

func (s *OffStateC) PressButton() {
	fmt.Println("Turning the light on...")
	s.lightBulb.SetState(s.lightBulb.onState) // Change to OnState
}

// LightBulb is the context class that switches between states
type LightBulbC struct {
	onState  CState
	offState CState
	state    CState
}

func NewLightBulbC() *LightBulbC {
	bulb := &LightBulbC{}
	onState := &OnStateC{lightBulb: bulb}
	offState := &OffStateC{lightBulb: bulb}

	bulb.onState = onState
	bulb.offState = offState
	bulb.state = offState // Initial state: off

	return bulb
}

func (l *LightBulbC) PressButton() {
	l.state.PressButton() // Delegate to the current state
}

func (l *LightBulbC) SetState(s CState) {
	l.state = s // Change the current state
}
