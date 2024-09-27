package main

import (
	"fmt"
	"go-design-patterns/pkg/prototype"
)

func main() {
	// Create original Circle and Rectangle
	originalCircle := &prototype.Circle{Radius: 5, Color: "red"}
	originalRectangle := &prototype.Rectangle{Width: 10, Height: 20, Color: "blue"}

	fmt.Println("Original Objects:")
	fmt.Println(originalCircle)
	fmt.Println(originalRectangle)

	// Clone the Circle and Rectangle
	clonedCircle := originalCircle.Clone().(*prototype.Circle)          // Type assert to Circle
	clonedRectangle := originalRectangle.Clone().(*prototype.Rectangle) // Type assert to Rectangle

	fmt.Println("\nCloned Objects:")
	fmt.Println(clonedCircle)
	fmt.Println(clonedRectangle)

	// Modify cloned object without affecting original
	clonedCircle.Color = "green"
	clonedRectangle.Width = 15

	fmt.Println("\nModified Cloned Objects:")
	fmt.Println(clonedCircle)
	fmt.Println(clonedRectangle)

	fmt.Println("\nOriginal Objects after cloning:")
	fmt.Println(originalCircle)
	fmt.Println(originalRectangle)
}
