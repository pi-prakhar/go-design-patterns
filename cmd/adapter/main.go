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
}
