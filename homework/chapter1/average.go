package main

import "fmt"

func main() {
	var arr [3]float64 = [3]float64{0.25, 1.25, 8.65}
	var sum float64 = 0
	for i := range arr {
		sum += arr[i]
	}
	fmt.Printf(fmt.Sprintf("Średnia wynosi: %.2f", sum/float64(len(arr))))
}
