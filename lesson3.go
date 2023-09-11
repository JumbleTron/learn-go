package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func findFile(fileName string, path string) (string, error) {
	var file string
	var findingError error

	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			findingError = err
			return err
		}
		if info.Name() == fileName {
			file = currentPath
			return nil
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if file == "" {
		return "", fmt.Errorf("Plik '%s' nie został znaleziony", fileName)
	}

	return file, findingError
}

func addNUmber(a, b int) int {
	return a + b
}

func calcualte(a, b int) (int, int, int) {
	sum := a + b
	diff := a - b
	multiple := a * b
	return sum, diff, multiple
}

func diffrentParams(a, b int) (int, interface{}) {
	sum := a + b

	if sum%2 == 0 {
		return sum, "Suma jest liczbą parzystą"
	} else {
		return sum, []string{"Suma jest liczbą nieparzystą", "To jest dodatkowy komunikat"}
	}
}

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func calc(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	fmt.Println("Suma: ", addNUmber(2, 3))
	fileName := "przykladowy.txt"
	path := "/sciezka/do/katalogu"

	foundedFile, err := findFile(fileName, path)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Znaleziono plik '%s' w ścieżce: %s\n", fileName, foundedFile)
	}

	sum, diff, multiple := calcualte(10, 5)

	fmt.Printf("Suma: %d\n", sum)
	fmt.Printf("Różnica: %d\n", diff)
	fmt.Printf("Iloczyn: %d\n", multiple)

	result, message := diffrentParams(10, 5)
	fmt.Printf("Wynik: %d\n", result)

	switch message.(type) {
	case string:
		fmt.Printf("Komunikat: %s\n", message)
	case []string:
		fmt.Printf("Komunikat: %v\n", message)
	}

	fmt.Printf("Suma liczb: %v\n", sumNumbers(10, 2, 3))
	fmt.Printf("Suma liczb: %v\n", sumNumbers(5))

	numbers := []int{1, 2, 3, 4}
	fmt.Printf("Suma liczb w tablicy: %v\n", sumNumbers(numbers...))

	sum2, diff2 := calc(10, 5)
	fmt.Printf("Suma: %d\n", sum2)
	fmt.Printf("Różnica: %d\n", diff2)
}
