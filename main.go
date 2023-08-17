package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	name := "John"
	builder.WriteString("Witaj, ")
	builder.WriteString(name)
	builder.WriteString(". Jak się masz?")
	fmt.Println(builder.String())

	name2 := "Mark"
	fmt.Println(fmt.Sprintf("Witaj, %s. Jak się masz?", name2))

	x := 0.1
	y := 0.2
	fmt.Println(x+y, x+y == 0.3)

	var a complex64 = complex(1, 2)
	var b complex64 = 2 + 4i
	var c complex64 = b - a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var someByte byte = 'A'
	var someRune rune = 'ą'

	fmt.Println(someByte)
	fmt.Println(someRune)

}
