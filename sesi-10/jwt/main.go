package main

import (
	"jwt/database"
	"jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
