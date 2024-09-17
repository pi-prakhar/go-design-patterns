package main

import (
	"fmt"
	"go-design-patterns/pkg/factory"
)

func main() {
	p := factory.NewPerson("jack", 38)
	p.Eyecount = 1
	fmt.Println(p)

	developerFactory := factory.NewEmployeeFactory("developer", 60000)
	managerFactory := factory.NewEmployeeFactory("Managaer", 80000)
	dev1 := developerFactory("Raj")
	manager1 := managerFactory("Rahul")
	fmt.Println(dev1)
	fmt.Println(manager1)

	engineerFactory := factory.NewEmployeeFactory2("engineer", 100000)
	eng1 := engineerFactory.Create("Parth")
	fmt.Println(eng1)
}
