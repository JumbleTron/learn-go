package main

import "fmt"

func main() {
	for i := 3; i <= 100; i++ {
		isPrimary := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrimary = false
			}
		}
		if isPrimary {
			fmt.Println(i)
		}
	}
}
