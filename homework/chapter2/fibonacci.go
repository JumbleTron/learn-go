package main

import "fmt"

func main() {
	var value int = 0
	var nextValue int = 1

	var itemsCount int

	fmt.Print("Podaj ile chcesz wypisać wyrazów ciągu fibonacciego: ")
	_, err := fmt.Scan(&itemsCount)
	if err != nil {
		fmt.Println("Wystąpił błąd podczas wczytywania danych:", err)
		return
	}

	for i := 0; i < itemsCount; i++ {
		fmt.Println(fmt.Sprintf("Element numer: %d to: %d", i+1, nextValue))
		nextValue += value
		value = nextValue - value
	}
}
