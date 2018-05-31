package routing_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Skrekkur/gophercon-gokub-ws/pkg/routing"
)

func TestBaseRouter(t *testing.T) {
	handler := routing.BaseRouter()
	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Error("Did not get correct statusCode %i expected %i", res.StatusCode, http.StatusOK)
	}
}
