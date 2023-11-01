package main

import (
	"log"

	"github.com/gartyom/wb-practice/L0/internal/app"
)

func main() {

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
