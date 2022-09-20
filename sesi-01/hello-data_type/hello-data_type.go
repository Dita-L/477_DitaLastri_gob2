package main

import "fmt"

func main() {
  //DATA TYPE INT DAN UINT
	var contohnegativ = -10
	fmt.Printf("%T\n", contohnegativ)
	var positif = 10
	fmt.Printf("%T\n", positif)
	//	hasilnya int kak, berarti uint yang ga bisa negatif. Di teks ppt kebalik, tapi di gambarnya bener.

	//FLOAT
	var desimal float64 = 36.73527

	fmt.Printf("default desimal = %f\n", desimal)
	fmt.Printf("4 karakter dengan 2 angka di belakang koma = %4.2f\n", desimal)

	// BOOL
	var kondisi bool = true
	fmt.Printf("is it permitted? %t\n", kondisi)

	//STRING
	var kalimat string = "diawali dan diakhiri tanda petik"
	fmt.Printf("Tipe Data %T, %s\n", kalimat, kalimat)
	var pesan string = `ini contoh "string" yang pake backticks :)`
	fmt.Println(pesan)

	//ZERO VALUE
	var kata string
	var benarsalah bool
	var nonDesimal int
	var Desimal uint
	fmt.Println(kata, benarsalah, nonDesimal, Desimal)

}
