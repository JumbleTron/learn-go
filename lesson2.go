package main

import (
	"fmt"
	"time"
)

func obsluzZapytanie1(kanal chan string) {
	// Symulacja obsługi zapytania
	time.Sleep(2 * time.Second)
	kanal <- "Odpowiedź na zapytanie 1"
}

func obsluzZapytanie2(kanal chan string) {
	// Symulacja obsługi zapytania
	time.Sleep(3 * time.Second)
	kanal <- "Odpowiedź na zapytanie 2"
}

func obsluzZapytanie3(kanal chan string) {
	// Symulacja obsługi zapytania
	time.Sleep(1 * time.Second)
	kanal <- "Odpowiedź na zapytanie 3"
}

func main() {
	kanal1 := make(chan string)
	kanal2 := make(chan string)
	kanal3 := make(chan string)

	// Rozpoczęcie obsługi zapytań w osobnych wątkach
	go obsluzZapytanie1(kanal1)
	go obsluzZapytanie2(kanal2)
	go obsluzZapytanie3(kanal3)

	// Wybór pierwszej dostępnej odpowiedzi
	select {
	case odpowiedz := <-kanal1:
		fmt.Println(odpowiedz)
	case odpowiedz := <-kanal2:
		fmt.Println(odpowiedz)
	case odpowiedz := <-kanal3:
		fmt.Println(odpowiedz)
	}

	userSelect := 1
	switch userSelect {
	case 1:
		fmt.Println("Wybrana odpowiedź jest niepoprawna")
	case 2:
		fmt.Println("Wybrana odpowiedź jest poprawna")
	default:
		fmt.Println("Niedozwolona opcja")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for j <= 2 {
		fmt.Println(fmt.Sprintf("For działajacy jak while: %d", j))
		j++
	}

	type User struct {
		id    int
		email string
	}
	var users [2]User = [2]User{
		{id: 1, email: "test@mail.com"},
		{id: 2, email: "test2@mail.com"},
	}
	for indeks, user := range users {
		fmt.Println(fmt.Sprintf("Na miejscu: %d, znajduje się uytkownik o id: %d", indeks, user.id))
	}
}
