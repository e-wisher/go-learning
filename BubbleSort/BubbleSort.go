package main

import (
	"errors"
	"fmt"
)

//BubbleSort sort list via algorithm
func BubbleSort(numbers []int) ([]int, error) {
	if len(numbers) == 0 {
		return []int{}, errors.New("Empty list")
	}

	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}

		}
	}
	return numbers, nil
}

func main() {
	n := []int{1, 5, 6, 3, 8, 2}
	numbers, err := BubbleSort(n)
	if err != nil {
		fmt.Println("Fatal")
	}
	fmt.Println(numbers)
}
