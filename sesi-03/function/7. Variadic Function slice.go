package main

import "fmt"

func sum(number ...int) int {
	total := 0

	for _, v := range number {
		total += v
	}
	return total
}

func main() {
	numberLists := []int{1, 2, 3, 4, 5}

	result := sum(numberLists...)

	fmt.Println("Result:", result)

}
