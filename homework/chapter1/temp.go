package main

import "fmt"

func main() {
	var tempInC int

	fmt.Print("Podaj temperaturę w stopniach Celsjusza: ")
	_, err := fmt.Scan(&tempInC)
	if err != nil {
		fmt.Println("Wystąpił błąd podczas wczytywania danych:", err)
		return
	}

	fmt.Println(fmt.Sprintf("%d °C = %.2f °F", tempInC, float64(tempInC)*1.8+32))
}
