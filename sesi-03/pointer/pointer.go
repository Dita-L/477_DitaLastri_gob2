package main

import "fmt"

func main() {
	var firstnumber int = 4
	var secondnumber *int = &firstnumber

	fmt.Println("firstnumber (val) :", firstnumber)
	fmt.Println("firstnumber (memory address) :", &firstnumber)

	fmt.Println("secondnumber (val) :", *secondnumber)
	fmt.Println("secondnumber (memory address) :", secondnumber)
}
