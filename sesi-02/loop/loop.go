package main

import "fmt"

func main() {

	/*
		//PENGGUNAAN "for"
		for i := 0; i < 3; i++ {
			fmt.Println("angka", i)
		}

		///Contoh 3 Penggunaan "for"
		var i = 0

		for {
			fmt.Println("angka", i)
			i++

			if i == 3 {
				break
			}
		}

		///Continue dan Break
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				continue
			}

			if i > 8 {
				break
			}

			fmt.Println("angka", i)
		}

	*/

	///Loop bersarang
	for a := 0; a < 5; a++ {
		fmt.Println("Perulangan - ", a+1)

		for j := 0; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println("\n")
	}

	///Loop berlabel
outerloop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan - ", i+1)

		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerloop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}

}
