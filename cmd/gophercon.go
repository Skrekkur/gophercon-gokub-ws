package main

import (
	"log"
	"net/http"

	"github.com/Skrekkur/gophercon-gokub-ws/pkg/routing"
)

func main() {

	log.Printf("Service is starting")
	r := routing.BaseRouter()

	http.ListenAndServe(":8000", r)
}
