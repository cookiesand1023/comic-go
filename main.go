package main

import (
	"log"
	"net/http"

	"github.com/comic-go/config"
)

func main() {
	config.SetRouter()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
