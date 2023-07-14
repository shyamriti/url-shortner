package main

import (
	"urlshortner/src/router"
)

func main() {
	// database.Connection()
	r := router.Router()
	r.Run(":8080")
}
