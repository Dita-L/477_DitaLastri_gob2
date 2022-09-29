package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"Airell", "Mailo", "Indah"}

	for _, v := range students {
		go introduce(v, c)
	}

	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)
}

func print(c <-chan string) {
	fmt.Println(<-c) // channel pada func print hanya menerima data
}

func introduce(student string, c chan<- string) {
	//channel pada func introduce hanya mengirim data
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}
