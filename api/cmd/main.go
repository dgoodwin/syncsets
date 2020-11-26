package main

import (
	"database/sql"
	"github.com/dgoodwin/syncsets/api/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("running syncsets-api")
	r := mux.NewRouter().StrictSlash(true)

	db, err := sql.Open("postgres", "user=postgres password=WYZVrmtdvuQlsq4hvo8C host=localhost dbname=syncsets sslmode=disable")
	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	clusterHandler := handlers.NewClusterHandler(db)
	r.HandleFunc("/clusters", clusterHandler.Get).Methods("GET")
	r.HandleFunc("/clusters", clusterHandler.Post).Methods("POST")

	syncsetHandler := handlers.NewSyncSetHandler(db)
	r.HandleFunc("/syncsets", syncsetHandler.Get).Methods("GET")
	r.HandleFunc("/syncsets", syncsetHandler.Post).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
