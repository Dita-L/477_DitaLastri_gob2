package main

import (
	"fmt"
	"strings"
)

func profile(nama string, favFoods ...string) {
	merge := strings.Join(favFoods, ", ")

	fmt.Println("Hello! I'm", nama)
	fmt.Println("I really like", merge)
}

func main() {
	profile("Dita", "Nasi", "Sate", "Ayam Bakar", "Ikan Bakar")

}
