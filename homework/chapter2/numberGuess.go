package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := rand.Intn(100)
	var inputDigit int
	for inputDigit != number {
		fmt.Print("Wprowadź liczbę: ")
		_, err := fmt.Scan(&inputDigit)
		if err != nil {
			fmt.Println("Wystąpił błąd podczas wczytywania danych:", err)
			return
		}
		if inputDigit > number {
			fmt.Println("Liczba jest mniejsza od: ", inputDigit)
		} else if inputDigit < number {
			fmt.Println("Liczba jest większa od: ", inputDigit)
		} else {
			fmt.Println("Zgadłeś")
		}
	}
}
