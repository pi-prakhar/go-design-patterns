package main

import (
	"fmt"
	"go-design-patterns/pkg/state"
)

func main() {
	vendingMachine := state.NewVendingMachine()

	err := vendingMachine.InsertMoney()
	if err != nil {
		fmt.Println(err)
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		fmt.Println(err)
	}

	lightBulbC := state.NewLightBulbC()

	lightBulbC.PressButton() // Turning the light on...
	lightBulbC.PressButton() // Turning the light off...

	lightBulbNC := state.NewLightBulbNC()

	lightBulbNC.PressButton() // Turning the light on...
	lightBulbNC.PressButton() // Turning the light off...

}
