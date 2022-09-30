package main

import "fmt"

func main() {
	callDeferFunc()
	fmt.Println("Hai everyone!")
}

func callDeferFunc() {
	fmt.Println("satu")
	defer deferFunc()
	fmt.Println("dua")
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
