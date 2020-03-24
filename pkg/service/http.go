package service

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func NewHTTPServer() *httprouter.Router {
	router := httprouter.New()
	router.HandleMethodNotAllowed = true

	router.Handle("GET", "/hostname", HostName)

	return router
}

func HostName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "hostname, %s!\n", hostname)
}
