package main

import (
	"fmt"
)

//BubbleSort sort list via algorithm
func BubbleSort(numbers []int) []int {

	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] < numbers[j] {
				skipholder := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = skipholder
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
