package factory

import "fmt"

type Person struct {
	Name     string
	Age      int
	Eyecount int
}

// creating a factory for person
func NewPerson(name string, age int) Person {
	//can add some validations
	if age < 16 {
		fmt.Printf("below age")
	}
	return Person{name, age, 2}
}
