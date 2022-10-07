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

	var temp interface{}

	var err = json.Unmarshal([]byte(jsonString), &temp)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result = temp.(map[string]interface{})

	fmt.Println("full_name", result["full_name"])
	fmt.Println("email:", result["email"])
	fmt.Println("age:", result["age"])
}
