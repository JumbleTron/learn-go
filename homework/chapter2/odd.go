package main

import "fmt"

func main() {
	var odd []int
	var even []int
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}
	fmt.Println(fmt.Sprintf("Liczby parzyste: %d", even))
	fmt.Println(fmt.Sprintf("Liczby nieparzyste: %d", odd))
}
