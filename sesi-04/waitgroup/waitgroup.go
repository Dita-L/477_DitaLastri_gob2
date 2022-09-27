package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"apple", "manggo", "durian", "rambutan"}

	var ab sync.WaitGroup
	//ab adalah variabel dengan tipe data sync.Waitgroup
	//sync.Waitgroup merupakan sebuah struct dari package sync

	//looping fruits
	for index, fruit := range fruits {
		ab.Add(1)
		//method Add() digunakan untuk menambah counter dari waitgroup
		//Counter dari waitgroup ini berguna untuk memberitahu waitgroup tentang jumlah go routine yang harus ditunggu.

		go printFruit(index, fruit, &ab)
	}
	//Karena kita melooping sebanyak 4 kali, berarti terdapat 4 go routine yang harus ditunggu sebelum function main menghentikan eksekusinya.

	ab.Wait()
	// Method Wait() berguna untuk menunggu seluruh go routine menyelesaikan proses nya,
	//Method Wait() akan menahan function main hingga seluruh proses go routine selesai.

}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	// Waitgroup pada function printFruit adalah sebuah pointer,
	//hal ini perlu perlu dilakukan agar waitgroup pada function main dan printFruit mengandung memori yang sama.

	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	//untuk memberitahu waitgroup tentang go routine yang telah menyelesaikan proses nya, maka kita perlu memanggil method Done()
	wg.Done()

}
