package main

import (
	"github.com/dgoodwin/syncsets/api/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("running syncsets-api")
	r := mux.NewRouter().StrictSlash(true)
	r.Handle("/clusters", handlers.NewClusterHandler())
	log.Fatal(http.ListenAndServe(":8080", r))
}
