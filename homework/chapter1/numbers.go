package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	//var input int
	sum := 0

	fmt.Print("Wprowadź liczbę: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Wystąpił błąd podczas wczytywania danych:", err)
		return
	}

	/*tempNumber := input
	for tempNumber != 0 {
		remainder := tempNumber % 10
		sum += remainder
		tempNumber /= 10
	}*/

	for _, char := range input {
		digitStr := string(char)
		digit, err := strconv.Atoi(digitStr)
		if err != nil {
			fmt.Printf("Błąd konwersji \"%s\" na liczbę: %s\n", digitStr, err)
			continue
		}
		sum += digit
	}

	fmt.Println(fmt.Sprintf("Suma liczb w: %s wynosi: %d", input, sum))
}
