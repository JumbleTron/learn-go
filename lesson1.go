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

	message := `To jest
	wieloliniowy
	string.`

	fmt.Println(message)

	var isTrue bool = true
	var isFalse bool = false

	fmt.Println(isTrue)
	fmt.Println(isFalse)

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

	data := []byte{65, 66, 67} // Reprezentacja ASCII znaków 'A', 'B', 'C'
	for _, b := range data {
		fmt.Printf("%c ", b) // Wypisze: A B C
	}

	text := "Hello, 世界" // Łańcuch zawierający znaki ASCII i znaki Unicode
	for _, r := range text {
		fmt.Printf("%c ", r) // Wypisze: H e l l o ,   世 界
	}

	fmt.Println(someByte)
	fmt.Println(someRune)

	var someComplex1 complex64 = complex(1, 2)
	var someComplex2 complex64 = 2 + 4i
	var someComplex3 complex64 = someComplex2 - someComplex1
	fmt.Println(someComplex1)
	fmt.Println(someComplex2)
	fmt.Println(someComplex3)

	var arr [3]int = [3]int{1, 2, 3}
	var slice []int = []int{4, 5, 6}

	fmt.Println(arr)
	fmt.Println(slice)

	var m map[string]int = map[string]int{"jabłko": 3, "banan": 5}
	fmt.Println(m)

	type Person struct {
		Name string
		Age  int
	}

	var p Person = Person{Name: "Jan", Age: 30}
	fmt.Println(p)

	someNumber1 := 10
	someNumber2 := 3

	sum := someNumber1 + someNumber2
	fmt.Println(fmt.Sprintf("Suma: %d", sum))
	difference := someNumber1 - someNumber2
	fmt.Println(fmt.Sprintf("Odejmowanie: %d", difference))
	product := someNumber1 * someNumber2
	fmt.Println(fmt.Sprintf("Mnozenie: %d", product))
	quotient := someNumber1 / someNumber2
	fmt.Println(fmt.Sprintf("Dzielenie: %d", quotient))
	remainder := someNumber1 % someNumber2
	fmt.Println(fmt.Sprintf("Reszta z dzielenia: %d", remainder))
}
