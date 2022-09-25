package main

import "fmt"

func changeVal(ptr *int) {
	*ptr = 20
}

func main() {
	var a int = 10
	fmt.Println("Before:", a)
	changeVal(&a)
	fmt.Println("After:", a)
}
