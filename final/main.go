package main

import (
	"final/database"
	"final/router"
)

func main() {
	database.GetDB()
	r := router.StartApp()
	r.Run(":8080")
}
