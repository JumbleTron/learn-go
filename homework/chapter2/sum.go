package main

import "fmt"

func main() {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf(fmt.Sprintf("Suma wynosi: %d", sum))
}
