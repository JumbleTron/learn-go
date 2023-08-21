package main

import (
	"fmt"
	"reflect"
)

func main() {
	matrix1 := [3][3]int{
		{0, 1, 2},
		{4, 5, 6},
		{8, 9, 10},
	}
	matrix2 := [3][3]int{
		{15, 3, 1},
		{69, 15, 2},
		{0, 1, 7},
	}
	var sum [3][3]int

	for i := range matrix1 {
		if len(matrix1) != len(matrix2) {
			fmt.Println("Macierze mają rózne długości")
			return
		}

		if reflect.TypeOf(matrix1[i]).Kind() == reflect.Array {
			if len(matrix1[i]) != len(matrix2[i]) {
				fmt.Println("Macierze mają rózne długości")
				return
			}
			for j := range matrix1[i] {
				if len(matrix1[i]) == len(matrix2[i]) {
					sum[i][j] = matrix2[i][j] + matrix1[i][j]
				}
			}
		}
	}
	fmt.Println(sum)
}
