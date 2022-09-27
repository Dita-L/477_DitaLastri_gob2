package main

import "fmt"

func main() {
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Airell", "Nanda"}

	fmt.Println(randomValues)
}
