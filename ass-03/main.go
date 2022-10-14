package main

import (
	"ass-03/router"
)

func main() {
	var PORT = ":8080"

	router.StartServer().Run(PORT)
}
