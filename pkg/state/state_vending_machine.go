package state

import "fmt"

// state interface -- also the strategy to be implemented by others
type VState interface {
	InsertMoney() error
	DispenseItem() error
}

// context
type VendingMachine struct {
	NoMoney  VState
	HasMoney VState
	state    VState
}

func (v *VendingMachine) InsertMoney() error {
	return v.GetState().InsertMoney()
}

func (v *VendingMachine) DispenseItem() error {
	return v.GetState().InsertMoney()
}

func (v *VendingMachine) GetState() VState {
	return v.state
}

func (v *VendingMachine) SetState(newState VState) {
	v.state = newState
}

func NewVendingMachine() *VendingMachine {
	vm := VendingMachine{}
	NoMoneyState := &NoMoneyState{VendingMachine: &vm}
	HasMoneyState := &HasMoneyState{VendingMachine: &vm}

	vm.HasMoney = HasMoneyState
	vm.NoMoney = NoMoneyState
	vm.state = NoMoneyState // initial state

	return &vm
}

// concrete impelementation 1
type NoMoneyState struct {
	VendingMachine *VendingMachine
}

func (s *NoMoneyState) InsertMoney() error {
	fmt.Println("Money Inserted!!")
	s.VendingMachine.SetState(s.VendingMachine.HasMoney)
	return nil
}

func (s *NoMoneyState) DispenseItem() error {
	return fmt.Errorf("Insert Money first!!")
}

// concrete implementation 2
type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (s *HasMoneyState) InsertMoney() error {
	return fmt.Errorf("Money already inserted!!")
}

func (s *HasMoneyState) DispenseItem() error {
	fmt.Println("Item Dispensed!!")
	s.VendingMachine.SetState(s.VendingMachine.NoMoney)
	return nil
}
