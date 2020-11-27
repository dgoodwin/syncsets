package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("running syncsets-controller")

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

	msgs, err := ch.Consume(
		"clusters", // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.WithError(err).Fatal("error consuming from queue")
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.WithField("body", string(d.Body)).Info("received message from queue")
		}
	}()

	log.Info("waiting for messages")
	<-forever
}
