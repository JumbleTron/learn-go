package main

import "fmt"

func main() {
	var arr []int = []int{3, 5, 100, 65, -3, 21, 698, 0}
	fmt.Println("Nieposortowana tablica", arr)
	fmt.Println("Posortowana tablica", insertionSort(arr))
}
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		currentValue := arr[i]
		prevIndex := i - 1
		for prevIndex >= 0 && arr[prevIndex] > currentValue {
			arr[prevIndex+1] = arr[prevIndex]
			prevIndex--
		}
		arr[prevIndex+1] = currentValue
	}

	return arr
}
