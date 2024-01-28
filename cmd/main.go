package main

import (
	"log"

	server "github.com/totoyk/trial-api-golang/internal"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
