package main

import (
	"github.com/wgeorgecook/testing-http/internal/pkg/api"
	"log"
	"net/http"
)

func main() {
	log.Print("Hello World!")
	defer log.Println("Goodbye!")

	a := api.Client{HttpClient: http.DefaultClient}
	resp, err := a.GetResourceByID("1")
	if err != nil {
		panic(err)
	}
	log.Print(resp)
}
