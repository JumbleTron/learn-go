package main

import "fmt"

func main() {
	var inputDigit int

	fmt.Print("Wprowadź dowolna liczbę: ")
	_, err := fmt.Scan(&inputDigit)
	if err != nil {
		fmt.Println("Wystąpił błąd podczas wczytywania danych:", err)
		return
	}
	digit := inputDigit

	var binaryString string
	var binaryStringRevert string

	if digit == 0 {
		binaryStringRevert = "0"
	}

	for digit != 0 {
		binaryString += fmt.Sprintf("%d", digit%2)
		digit = digit / 2
	}
	for i := len(binaryString) - 1; i >= 0; i-- {
		binaryStringRevert += string(binaryString[i])
	}
	fmt.Println(fmt.Sprintf("Liczba: %d w zapisie binarnym to: %s", inputDigit, binaryStringRevert))
}
