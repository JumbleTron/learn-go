package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := "To jest przykładowy tekst. Jestem pewjsien, że jestem w stanie to zrobić w js oraz go"
	searchString := "js"

	words := strings.Fields(strings.ToLower(text))
	for i, word := range words {
		if strings.ToLower(word) == searchString {
			fmt.Println(fmt.Sprintf("Pozycja %d: %s", i+1, word))
		}
	}

	regex := regexp.MustCompile("js")
	matches := regex.FindAllStringIndex(strings.ToLower(text), -1)

	for _, match := range matches {
		start := match[0]
		end := match[1]
		fmt.Println(fmt.Sprintf("Pozycja %d-%d: %s", start, end, text[start:end]))
	}

}
