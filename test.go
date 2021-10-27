package main

import (
	"io"
	"log"
	"net/http"

	api "github.com/argSea/nauplius"
)

func main() {
	c := api.Controller{}
	a := api.API{Controller: c}
	err := http.ListenAndServe(":8080", &a)

	rWriter := a.GetHTTPWriter()
	test := a.GetData()

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	} else {
		io.WriteString(rWriter, test)
	}
}
