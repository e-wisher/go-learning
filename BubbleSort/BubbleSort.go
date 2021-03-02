package main

import (
	"fmt"
)

//BubbleSort sort list via algorithm
func BubbleSort(numbers []int) []int {

	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}

		}
	}
	fmt.Println(numbers)
	return numbers
}

func main() {
	numberslist := []int{1, 5, 6, 3, 2}
	BubbleSort(numberslist)
}
