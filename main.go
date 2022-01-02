package main

import (
	"log"

	"github.com/yescorihuela/beers_app/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
