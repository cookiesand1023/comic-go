package main

import (
	"log"
	"net/http"
	"os"

	"github.com/comic-go/config"
	"github.com/comic-go/model"
)

func main() {
	config.SetRouter()
	config.LoadEnv()
	model.InitDatabase()

	log.Printf("this environment is " + os.Getenv("GO_ENV"))
	log.Printf("start go server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
