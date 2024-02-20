package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	r := &Rectangle{Width: 5, Height: 10}
	PrintArea(r)
	num := 7
	num2 := 20
	fmt.Printf("Number with sign: %+d\n", num2)         // Wyświetli "+20"
	fmt.Printf("Number without sign: %d\n", num2)       // Wyświetli "20"
	fmt.Printf("Number without sign: %d\n", -num2)      // Wyświetli "-20"
	fmt.Printf("Number with sign: %+d\n", -num2)        // Wyświetli "-20"
	fmt.Printf("Number filled with zeros: %05d\n", num) // Wyświetli "00007"
	name := "John"
	fmt.Printf("Name aligned to the left: %-10s\n", name) // Wyświetli "John      "
}
