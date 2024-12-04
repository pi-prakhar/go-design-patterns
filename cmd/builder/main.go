package main

import (
	"fmt"
	"go-design-patterns/pkg/builder"
	"strings"
)

func main() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("Bye")
	sb.WriteString(hello)
	sb.WriteString("Bye")
	fmt.Println(sb.String())

	words := []string{"hello", "world"}
	sb.Reset()
	sb.WriteString(words[0])
	sb.WriteString(words[1])
	fmt.Println(sb.String())

	bb := builder.NewBurgerBuilder()
	burger := bb.ChooseBread(builder.MultiGrain).ChoosePatty(builder.ChickenPatty).AddCheese().AddMayonaise().Build()
	fmt.Println(burger)
}
