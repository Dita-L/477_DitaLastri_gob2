package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Airell", Age: 23, Division: "IT"},
	{ID: 2, Name: "Nanda", Age: 23, Division: "Finance"},
	{ID: 3, Name: "Mailo", Age: 20, Division: "IT"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)

}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//method header dari interface http.ResponseWriter, di chain dengan method Set()
	//hal ini dilakukan untuk menentukan bentuk dari data response yang ingin kita kirimkan ke client
	//karena kita ingin mengirim response dalam bentuk json, maka content type menjadi application/json dalam method Set

	if r.Method == "GET" {
		//lakukan pengecekan method
		//function getEmployees kita gunakan untuk mendapatkan data dari server menggunakan method GET

		json.NewEncoder(w).Encode(employees)
		//mengkonversi data employees menjadi data berbentuk json untuk dikirimkan kepada client
		//menggunakan method NewEncoder dari package json, di chain dengan method Encode untuk mengkonversi datanya menjadi bentuk json
		return
	}

	http.Error(w, "invalid method", http.StatusBadRequest)
	//jika method yang dikirimkan client bukan GET,
	//kita akan mengirimkan response error dengan menggunakan function Error dari package http
	//http.StatusBadRequest merupakan sebuah konstanta dari package http.StatusBadRequest,
	//merepresentasikan status 400

}
