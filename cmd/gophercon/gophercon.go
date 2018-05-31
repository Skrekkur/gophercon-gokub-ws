package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Skrekkur/gophercon-gokub-ws/pkg/routing"
	"github.com/Skrekkur/gophercon-gokub-ws/pkg/webserver"
)

func startServer(port string, router http.Handler, friendlyName string) *webserver.WebServer {

	if len(port) == 0 {
		log.Fatal("Service port wasn't set for" + friendlyName)
	}

	log.Printf(friendlyName + " Service is starting on port:" + port)
	ws := webserver.New("", port, router)
	go func() {
		log.Fatal(ws.Start())
	}()
	return ws
}
func main() {
	//$env:SERVICE_PORT = 1000

	shutdown := make(chan error, 2)
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	port := os.Getenv("SERVICE_PORT")
	internalPort := os.Getenv("INTERNAL_PORT")

	normalRouter := routing.BaseRouter()

	diagnosticsRouter := routing.DiagnosticsRouter()
	ws := startServer(port, normalRouter, "normalServer")
	internalServer := startServer(internalPort, diagnosticsRouter, "healthService")

	select {
	case killSignal := <-interrupt:
		log.Printf(" Getting killed %s", killSignal)

	case err := <-shutdown:
		log.Printf("Got an error shutting down%s", err)
	}

	log.Printf("Stopping Services")
	ws.Stop()
	internalServer.Stop()
	os.Exit(0)
}
