package main

import (
	"log"

	"github.com/onunez-g/jsonresume-api/routes"
)

func main() {
	r := routes.GetRoutes()
	log.Fatal(r.Run())
}
