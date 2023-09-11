package main

import "fmt"

func main() {
	var inputDigit int
	sum := 0

	fmt.Print("Wprowadź liczbę: ")
	_, err := fmt.Scan(&inputDigit)
	if err != nil {
		fmt.Println("Wystąpił błąd podczas wczytywania danych:", err)
		return
	}

	for i := 1; i <= inputDigit; i++ {
		sum += i
	}

	fmt.Println(fmt.Sprintf("Suma liczb od: 1 do: %d to: %d", inputDigit, sum))

}
