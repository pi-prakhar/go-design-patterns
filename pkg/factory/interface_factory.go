package factory

import "fmt"

type PersonI interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s \n", p.name)
}

func NewPersonI(name string, age int) PersonI {
	return &person{name: name, age: age}
}
