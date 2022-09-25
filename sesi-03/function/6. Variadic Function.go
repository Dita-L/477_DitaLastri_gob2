package main

import "fmt"

func main() {
	studentLists := print("ariel", "nanda", "nailo", "marco")

	fmt.Printf("%v", studentLists)

}

func print(names ...string) []map[string]string {
	//fungsi "print" ini adalah "variadic function"
	//yaitu, func yang dapat menerima banyak argumen.
	//variadic functions ditandai adanya titik 3x (...) sebelum tipedata yang diterima.

	var result []map[string]string //map kosong

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		//key berperan sebagai key pada map
		temp := map[string]string{
			key: v,
		}
		//v sebagai value, temp sebagai map
		result = append(result, temp)
		//result = kosong + temp
	}
	return result
}
