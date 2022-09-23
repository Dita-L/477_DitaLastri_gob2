package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//
	city := "jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}

	//
	var kota []byte = []byte{106, 97, 107, 97, 114, 116, 97}
	fmt.Println(kota)
	fmt.Println(string(kota))

	//
	str1 := "makan"
	str2 := "mânca"

	fmt.Printf("str1 byte length = %d\n", len(str1))
	fmt.Printf("str2 byte length = %d\n", len(str2))

	fmt.Printf("str2 byte length = %d\n", utf8.RuneCountInString(str2))

	test := utf8.RuneCountInString(str2)
	_ = test

	for i, v := range str2 {
		fmt.Printf("indeks: %d, chara: %s\n", i, string(v))
	}

}
