package main

import (
	"log"
	"os"

	"github.com/Skrekkur/gophercon-gokub-ws/pkg/routing"
	"github.com/Skrekkur/gophercon-gokub-ws/pkg/webserver"
)

func main() {
	//$env:SERVICE_PORT = 1000
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port wasn't set")
	}
	log.Printf("Service is starting on port:" + port)
	r := routing.BaseRouter()

	ws := webserver.New("", port, r)
	ws.Start()
	//log.Fatal(http.ListenAndServe(":"+port, r))
}
