package main

import (
	"fmt"
	"go-design-patterns/pkg/factory"
)

func main() {
	p := factory.NewPerson("jack", 38)
	p.Eyecount = 1

	fmt.Println(p)
}
