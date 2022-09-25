package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15

	var area, circumference = calculate(diameter)

	fmt.Println("Area:", area)
	fmt.Println("circumference:", circumference)
}

func calculate(d float64) (area float64, circumference float64) {
	//luas
	area = math.Pi * math.Pow(d/2, 2)
	//keliling
	circumference = math.Pi * d

	return area, circumference
}
