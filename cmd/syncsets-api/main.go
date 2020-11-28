package main

import (
	"database/sql"
	"github.com/dgoodwin/syncsets/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"net/http"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("running syncsets-api")
	r := mux.NewRouter().StrictSlash(true)

	conn, err := amqp.Dial("amqp://ffY5PQ_tMehsn2tryfCDvuEvVDIBoLYu:Sp7vDEG8J_E62XFi-6r3XWBQJJi0T1Sy@rabbitmq:5672/")
	if err != nil {
		log.WithError(err).Fatal("error connecting to to rabbitmq")
	}
	defer conn.Close()
	log.Info("rabbitmq connection established")
	ch, err := conn.Channel()
	if err != nil {
		log.WithError(err).Fatal("error establishing rabbitmq channel")
	}
	defer ch.Close()

	db, err := sql.Open("postgres", "user=postgres password=WYZVrmtdvuQlsq4hvo8C host=postgresql dbname=syncsets sslmode=disable")
	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	clusterHandler := handlers.NewClusterHandler(db)
	r.HandleFunc("/clusters", clusterHandler.Get).Methods("GET")
	r.HandleFunc("/clusters", clusterHandler.Post).Methods("POST")
	q, err := ch.QueueDeclare(
		"clusters", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.WithError(err).Fatal("error declaring rabbitmq queue")
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("hello world"),
		})
	if err != nil {
		log.WithError(err).Error("error publishing test message")
	}

	syncsetHandler := handlers.NewSyncSetHandler(db)
	r.HandleFunc("/syncsets", syncsetHandler.Get).Methods("GET")
	r.HandleFunc("/syncsets", syncsetHandler.Post).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
