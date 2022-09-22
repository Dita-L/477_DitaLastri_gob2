package main

import "fmt"

func main() {
	/*	//LOOPING
		//FIRST WAY

		for i := 0; i <= 5; i++ {
			fmt.Println("Angka", i)
		}

		//Second way ====== variabel di sebut diawal

		var a = 3

		for a < 5 {
			fmt.Println(a)
			a++
		}

		//Third Way

		var x = 7

		for {
			fmt.Println(x)
			x--

			if x == 4 {
				break
			}
		}
	*/

	//Array
	var banyak_angka [4]int
	banyak_angka = [4]int{4, 3, 2, 1}
	fmt.Println(banyak_angka)

	var angka_angka = [4]int{4, 3, 2, 1}
	fmt.Println(angka_angka)

	kata_kata := [3]string{"halo", "hai", "bye"}
	fmt.Println(kata_kata)

	var buah = [4]string{"mangga", "apel", "jambu", "jeruk"}
	fmt.Println(buah)
	buah[0] = "rambutan"
	fmt.Println(buah)

	//looping pada array menggunakan "range"
	for i, element := range buah {
		fmt.Printf("index: %d, elemen: %s\n", i, element)
	}

	//looping pada arrary menggunakan "len"
	for i := 0; i < len(buah); i++ {
		fmt.Printf("index= %d, elemen: %s\n", i, buah[i])
	}

	//Multidimensional Array
	var contoh [3][4]int
	fmt.Println(contoh)

}
