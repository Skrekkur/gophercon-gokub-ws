package routing

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// ReadynessChecks is a sample
func DiagnosticsRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/readyz", readyzHandler()).Methods(http.MethodGet)
	r.HandleFunc("/healthz", readyzHandler()).Methods(http.MethodGet)
	return r
}

func readyzHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusText(http.StatusOK))
	}

}
