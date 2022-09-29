package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"Airell", "Mailo", "Indah"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c <- result
		}(v) //langsung memanggil func dengan argumen v, yaitu nama-nama students

	}

	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)
}

func print(c chan string) {
	fmt.Println(<-c)
}
