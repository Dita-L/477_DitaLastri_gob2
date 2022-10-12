package main

import (
	"jsonToken/database"
	"jsonToken/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
