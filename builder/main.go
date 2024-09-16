package main

import (
	"fmt"
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
}
