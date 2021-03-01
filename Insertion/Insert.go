package main

import (
	"fmt"
)

func InsertionSort(numbers []int) []int {

	for i := 1; i < len(numbers); i++ {
		for j := 0; j < i; j++ {
			if numbers[j] > numbers[i] {
				numbers[j], numbers[i] = numbers[i], numbers[j]
			}
		}
	}
	return numbers
}

func main() {
	result := InsertionSort([]int{2, 7, 6, 3, 4})
	fmt.Println(result)
}
