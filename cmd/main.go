package main

import (
	"log"

	server "github.com/totoyk/sample-backendapp-golang-on-echo/internal"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
