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

func someFunc(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func generateNextNumber() func() int {
	num := 0
	increess := func() int {
		num++
		return num
	}

	return increess
}

type Person struct {
	firstName string
	age       int
}

// Przekazywanie typu wartości
func changeFirstName(currentPerson Person, firstName string) {
	currentPerson.firstName = firstName
}

// Przekazywanie wskaźnika do typu wartości
func changeAge(currentPerson *Person, age int) {
	currentPerson.age = age
}
func sumNumbersNoPF(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
		numbers[i] += 10
	}
	return sum
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

	addanonymousFunc := func(a, b int) int {
		return a + b
	}

	sumOfNumbers := addanonymousFunc(3, 5)
	fmt.Printf("Suma: %d\n", sumOfNumbers)

	resultWith := someFunc(10, 4, func(x, y int) int {
		return x * y
	})

	fmt.Printf("Wynik operacji: %d\n", resultWith)

	numberGenerator := generateNextNumber()

	fmt.Println(numberGenerator())
	fmt.Println(numberGenerator())
	fmt.Println(numberGenerator())

	person := Person{firstName: "Anna", age: 30}

	changeFirstName(person, "Alice")
	fmt.Println("Imie po zmianie imienia:", person.firstName) // Nadal "Anna", bo osoba jest niezmienne

	changeAge(&person, 35)
	fmt.Println("Wiek po zmianie wieku:", person.age) // Teraz 35, ponieważ zmieniliśmy oryginalną osobę

	var numbersSlcie []int = []int{5, 10}
	fmt.Println(sumNumbersNoPF(numbersSlcie))
	fmt.Println("Wartość zmiennej numbersSlcie: ", numbersSlcie)

}
