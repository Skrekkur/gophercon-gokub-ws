package webserver

import "net"
import "net/http"

type WebServer struct {
	http.Server
	address string
}

//Constructs a new Webserver
func New(host, port string, h http.Handler) *WebServer {
	var ws WebServer

	ws.Addr = net.JoinHostPort(host, port)
	ws.Handler = h
	return &ws
}

func (s *WebServer) Start() error {
	return s.ListenAndServe()
}

func (s *WebServer) Stop() error {
	return s.Stop()
}
