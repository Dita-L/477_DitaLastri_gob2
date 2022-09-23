package main

import "fmt"

func main() {

	//MAP
	//membuat map

	var person map[string]string //Deklarasi
	person = map[string]string{
		"nama":     "Dita Lastri",
		"usia":     "24",
		"kotaasal": "Kab.Bekasi",
	}
	fmt.Println(person)

	//menambahkan key:value pada map
	person["hobby"] = "jalan"
	fmt.Println(person)

	//mengakses dan mengubah data
	person["nama"] = "Dita" //mengubah

	fmt.Println(person)
	fmt.Println("saya berasal dari", person["kotaasal"]) //mengakses

	//menghapus data
	delete(person, "hobby")
	fmt.Println(person)

	//mendeteksi ada tidaknya data
	value, adatidak := person["hobby"]
	fmt.Println(value, adatidak)

	//loop map
	for key, val := range person {
		fmt.Println(key, ":", val)
	}

}
