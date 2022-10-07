package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `
		{
			"full_name": "Airell Jordan",
			"email": "airell@mail.com",
			"age": 23
		}
	`

	var result Employee

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name", result.FullName)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)
}
