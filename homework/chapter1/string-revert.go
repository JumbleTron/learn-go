package main

import "fmt"

func main() {
	someString := "simple"
	var revertString string

	index := len(someString) - 1
	for index >= 0 {
		revertString += string(someString[index])
		index--
	}

	fmt.Println(fmt.Sprintf("Reverted string: %s is: %s", someString, revertString))
}
