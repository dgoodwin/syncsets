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

	clusterHandler := handlers.NewClusterHandler()
	r.HandleFunc("/clusters", clusterHandler.Get).Methods("GET")
	r.HandleFunc("/clusters", clusterHandler.Post).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
