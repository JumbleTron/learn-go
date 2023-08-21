package main

import "fmt"

func main() {
	var input string

	fmt.Print("Wprowadź ciąg znaków: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Wystąpił błąd podczas wczytywania danych:", err)
		return
	}

	var revertString string

	index := len(input) - 1
	for index >= 0 {
		revertString += string(input[index])
		index--
	}

	if input != revertString {
		fmt.Println(fmt.Sprintf("Wprowadzony ciąg znaków: %s nie jest palidromem", input))
	} else {
		fmt.Println(fmt.Sprintf("Wprowadzony ciąg znaków: %s jest palidromem", input))
	}
}
