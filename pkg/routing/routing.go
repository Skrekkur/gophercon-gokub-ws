package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// BaseRouter is a sample
func BaseRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler2()).Methods(http.MethodGet)
	return r
}

func homeHandler2() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is processing: %s", r.URL.Path)
		fmt.Fprint(w, "Request processed in handler2")
	}

}
