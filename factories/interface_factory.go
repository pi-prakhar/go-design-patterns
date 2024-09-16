package factories

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Println("Hi, my name is %s", p.name)
}

func NewPerson(name string, age int) Person {
	return &person{name: name, age: age}
}

func main() {
	p := NewPerson("James", 34)
	p.SayHello()
}
