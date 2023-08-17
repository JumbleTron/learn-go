package main

import (
	"fmt"
	"strings"
)

func main() {
	someString := "fgegthtyte45tgrdft345"

	for pos, char := range someString {
		if strings.Count(someString, string(char)) == 1 {
			fmt.Println(fmt.Sprintf("Pierwszy niepowtórzony znak to: %s, znajuje się na pozycji: %d", string(char), pos+1))
			return
		}
	}

	fmt.Println("Brak niepowtarzalnych znaków.")
}
