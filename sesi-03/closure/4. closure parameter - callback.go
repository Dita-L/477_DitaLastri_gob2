package main

import (
	"fmt"
)

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddnumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddnumbers++
		}
	}
	return totalOddnumbers
}

func main() {
	var numbers = []int{2, 3, 8, 10, 3, 99, 23}

	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find)
}
