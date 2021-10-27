package main

import (
	"io"
	"log"
	"net/http"

	api "github.com/argSea/nauplius"
)

func main() {
	c := api.Controller{}
	e := api.APIEndpoint{}

	e.AddNewEndpoint("/test", handleTest)

	a := api.Router{Controller: c}
	err := http.ListenAndServe(":8080", &a)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func handleTest(w http.ResponseWriter) {
	io.WriteString(w, "Test")
}
