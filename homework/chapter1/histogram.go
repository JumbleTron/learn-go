package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Wprowadź ciąg znaków: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Wystąpił błąd podczas wczytywania danych:", err)
		return
	}

	lowerString := strings.ToLower(input)
	var histogram map[string]int = make(map[string]int)

	for i := 0; i < len(lowerString); i++ {
		character := string(lowerString[i])
		_, ok := histogram[character]
		if !ok {
			histogram[character] = 1
		} else {
			histogram[character] += 1
		}
	}

	for letter, occured := range histogram {
		fmt.Println("Litera:", letter, "występuję:", occured, "razy")
	}
}
