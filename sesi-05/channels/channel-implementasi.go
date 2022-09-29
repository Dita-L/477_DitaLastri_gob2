package main

import "fmt"

func main() {
	c := make(chan string)

	go introduce("Airell", c)

	go introduce("Nanda", c)

	go introduce("Mailo", c)

	//kita membuat goroutine sebanyak 3x

	//psn1 <- c yang artinya adalah function main menerima data dari function introduce dan disimpan pada variable psn1
	psn1 := <-c
	fmt.Println(psn1)

	psn2 := <-c
	fmt.Println(psn2)

	psn3 := <-c
	fmt.Println(psn3)

	close(c)
	/*close merupakan sebuah function yang digunakan untuk mengindikasikan bahwa sebuah channel
	sudah tidak digunakan untuk berkomunikasi lagi*/
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)

	c <- result
}
