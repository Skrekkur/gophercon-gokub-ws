package main

import (
	"log"
	"os"

	"github.com/Skrekkur/gophercon-gokub-ws/pkg/routing"
	"github.com/Skrekkur/gophercon-gokub-ws/pkg/webserver"
)

func main() {
	//$env:SERVICE_PORT = 1000

	shutdown := make(chan error, 2)
	interrupt := make(chan os.Signal, 1)

	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port wasn't set")
	}
	log.Printf("Service is starting on port:" + port)
	r := routing.BaseRouter()

	ws := webserver.New("", port, r)
	go func() {
		log.Fatal(ws.Start())
	}()

	internalPort := os.Getenv("INTERNAL_PORT")
	if len(internalPort) == 0 {
		log.Fatal("Internal port wasn't set")
	}

	diagnosticsRouter := routing.DiagnosticsRouter()
	diagnosticsServer := webserver.New(
		"", internalPort, diagnosticsRouter,
	)
	log.Fatal(diagnosticsServer.Start())

	select {
	case killSignal := <-interrupt:
		log.Printf(" Getting gilled %s", killSignal)

	case err := <-shutdown:
		log.Printf("Got an error shutting down%s", err)
	}

	ws.Stop()
	diagnosticsServer.Stop()
}
