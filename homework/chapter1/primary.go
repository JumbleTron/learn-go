package main

import "fmt"

func main() {
	var input int

	fmt.Print("Wprowadź liczbę: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Wystąpił błąd podczas wczytywania danych:", err)
		return
	}
	if input < 2 {
		fmt.Println(fmt.Sprintf("Liczb: %d nie jest liczbą pierwszą", input))
		return
	}

	for i := 2; i*i <= input; i++ {
		if input%i == 0 {
			fmt.Println(fmt.Sprintf("Liczb: %d nie jest liczbą pierwszą", input))
			return
		}
	}

	fmt.Println(fmt.Sprintf("Liczb: %d jest liczbą pierwszą", input))
}
