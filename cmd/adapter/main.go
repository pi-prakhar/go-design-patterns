package main

import (
	"fmt"
	"go-design-patterns/pkg/adapter"
)

func main() {
	rc := adapter.NewRectangle(6, 4)
	a := adapter.VectorToRaster(rc) // adapter!
	_ = adapter.VectorToRaster(rc)  // adapter!
	fmt.Print(adapter.DrawPoints(a))

	//adapter2 code
	applePhone := adapter.ApplePhone{}
	androidPhone := adapter.AndroidPhone{}
	windowsPhone := adapter.WindowsPhone{IsFastChargeSupported: true}
	windowsPhoneAdapter := adapter.WindowsPhoneAdapter{WindowsPhone: &windowsPhone}

	client := adapter.Client{}
	client.Charge(&androidPhone)
	client.Charge(&applePhone)
	// client.Charge(&windowsPhone)
	client.Charge(&windowsPhoneAdapter)
}
