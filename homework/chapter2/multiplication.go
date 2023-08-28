package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			divider := "   | "
			result := i * j
			if (j >= 2 || j <= 10) && result >= 10 {
				divider = "  | "
			}
			if i == 10 {
				divider = " | "
			}
			if result == 100 {
				divider = "|"
			}
			fmt.Print(fmt.Sprintf("%d x %d: %d %s", i, j, result, divider))
		}
		fmt.Print("\n")
	}
}
